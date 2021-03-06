// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package tchannel

import (
	"fmt"
	"sync"

	"go.uber.org/yarpc/api/peer"
	"go.uber.org/yarpc/api/transport"
	intsync "go.uber.org/yarpc/internal/sync"
	"go.uber.org/yarpc/peer/hostport"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/tchannel-go"
)

// Transport is a TChannel transport suitable for use with YARPC's peer
// selection system.
// The transport implements peer.Transport so multiple peer.List
// implementations can retain and release shared peers.
// The transport implements transport.Transport so it is suitable for lifecycle
// management.
type Transport struct {
	lock sync.Mutex
	once intsync.LifecycleOnce

	ch     Channel
	router transport.Router
	tracer opentracing.Tracer
	name   string
	addr   string

	peers map[string]*hostport.Peer
}

// NewTransport is a YARPC transport that facilitates sending and receiving
// YARPC requests through TChannel.
// It uses a shared TChannel Channel for both, incoming and outgoing requests,
// ensuring reuse of connections and other resources.
//
// Either the local service name (with the ServiceName option) or a user-owned
// TChannel (with the WithChannel option) MUST be specified.
func NewTransport(opts ...TransportOption) (*Transport, error) {
	var options transportOptions
	options.tracer = opentracing.GlobalTracer()
	for _, opt := range opts {
		opt(&options)
	}

	if options.ch != nil {
		return nil, fmt.Errorf("NewTransport does not accept WithChannel, use NewChannelTransport")
	}

	return options.newTransport(), nil
}

func (o transportOptions) newTransport() *Transport {
	return &Transport{
		once:   intsync.Once(),
		name:   o.name,
		addr:   o.addr,
		tracer: o.tracer,
		peers:  make(map[string]*hostport.Peer),
	}
}

// ListenAddr exposes the listen address of the transport.
func (t *Transport) ListenAddr() string {
	return t.addr
}

// RetainPeer adds a peer subscriber (typically a peer chooser) and causes the
// transport to maintain persistent connections with that peer.
func (t *Transport) RetainPeer(pid peer.Identifier, sub peer.Subscriber) (peer.Peer, error) {
	t.lock.Lock()
	defer t.lock.Unlock()

	hppid, ok := pid.(hostport.PeerIdentifier)
	if !ok {
		return nil, peer.ErrInvalidPeerType{
			ExpectedType:   "hostport.PeerIdentifier",
			PeerIdentifier: pid,
		}
	}

	p := t.getOrCreatePeer(hppid)
	p.Subscribe(sub)
	return p, nil
}

// **NOTE** should only be called while the lock write mutex is acquired
func (t *Transport) getOrCreatePeer(pid hostport.PeerIdentifier) *hostport.Peer {
	if p, ok := t.peers[pid.Identifier()]; ok {
		return p
	}

	p := hostport.NewPeer(pid, t)
	p.SetStatus(peer.Available)

	t.peers[p.Identifier()] = p

	return p
}

// ReleasePeer releases a peer from the peer.Subscriber and removes that peer
// from the Transport if nothing is listening to it.
func (t *Transport) ReleasePeer(pid peer.Identifier, sub peer.Subscriber) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	p, ok := t.peers[pid.Identifier()]
	if !ok {
		return peer.ErrTransportHasNoReferenceToPeer{
			TransportName:  "tchannel.Transport",
			PeerIdentifier: pid.Identifier(),
		}
	}

	if err := p.Unsubscribe(sub); err != nil {
		return err
	}

	if p.NumSubscribers() == 0 {
		delete(t.peers, pid.Identifier())
	}

	return nil
}

// Start starts the TChannel transport. This starts making connections and
// accepting inbound requests. All inbounds must have been assigned a router
// to accept inbound requests before this is called.
func (t *Transport) Start() error {
	return t.once.Start(t.start)
}

func (t *Transport) start() error {
	chopts := tchannel.ChannelOptions{
		Tracer: t.tracer,
		Handler: handler{
			router: t.router,
			tracer: t.tracer,
		},
	}
	ch, err := tchannel.NewChannel(t.name, &chopts)
	if err != nil {
		return err
	}
	t.ch = ch

	// Default to ListenIP if addr wasn't given.
	addr := t.addr
	if addr == "" {
		listenIP, err := tchannel.ListenIP()
		if err != nil {
			return err
		}

		addr = listenIP.String() + ":0"
		// TODO(abg): Find a way to export this to users
	}

	// TODO(abg): If addr was just the port (":4040"), we want to use
	// ListenIP() + ":4040" rather than just ":4040".

	if err := t.ch.ListenAndServe(addr); err != nil {
		return err
	}

	t.addr = t.ch.PeerInfo().HostPort

	return nil
}

// Stop stops the TChannel transport. It starts rejecting incoming requests
// and draining connections before closing them.
// In a future version of YARPC, Stop will block until the underlying channel
// has closed completely.
func (t *Transport) Stop() error {
	return t.once.Stop(t.stop)
}

func (t *Transport) stop() error {
	t.ch.Close()
	return nil
}

// IsRunning returns whether the TChannel transport is running.
func (t *Transport) IsRunning() bool {
	return t.once.IsRunning()
}
