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

package retry

import (
	"context"
	"time"

	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/internal/backoff"
)

// PolicyProvider returns a retry policy to use for the request,
// if no policy is available it should return a default policy.
type PolicyProvider func(context.Context, *transport.Request) *Policy

var defaultPolicy = Policy{
	retries: 0,
	timeout: time.Second,
	backoff: backoff.FixedBackoff(time.Millisecond * 200).Backoff,
}

// Policy defines how a retry will be applied.  It contains all the information
// needed to preform a retry.
type Policy struct {
	// retries is the number of attempts we will retry (after the
	// initial attempt.
	retries uint

	// timeout is the Timeout we will enforce per request (if this
	// is less than the context deadline, we'll use that instead).
	timeout time.Duration

	// backoff is a backoff strategy that will be called after every retry.
	backoff backoff.Strategy
}

func (p *Policy) getTimeout(ctx context.Context) time.Duration {
	ctxDeadline, ok := ctx.Deadline()
	if !ok {
		return p.timeout
	}
	now := time.Now()
	if ctxDeadline.After(now.Add(p.timeout)) {
		return p.timeout
	}
	return ctxDeadline.Sub(now)
}

func (p *Policy) getBackoff(attempt uint) time.Duration {
	if p.backoff == nil {
		return time.Duration(0)
	}

	return p.backoff(attempt)
}

// NewPolicy creates a new retry Policy that can be used in retry middleware.
func NewPolicy(opts ...PolicyOption) *Policy {
	policy := defaultPolicy
	for _, opt := range opts {
		opt(&policy)
	}
	return &policy
}

// PolicyOption customizes the behavior of a retry policy.
type PolicyOption func(*Policy)

// Retries is the number of attempts we will retry (after the
// initial attempt.
//
// Defaults to 1.
func Retries(retries uint) PolicyOption {
	return func(pol *Policy) {
		pol.retries = retries
	}
}

// PerRequestTimeout is the Timeout we will enforce per request (if this
// is less than the context deadline, we'll use that instead).
//
// Defaults to 1 second.
func PerRequestTimeout(timeout time.Duration) PolicyOption {
	return func(pol *Policy) {
		pol.timeout = timeout
	}
}

// BackoffStrategy sets the backoff strategy that will be used after each
// failed request.
//
// Defaults to a 200ms backoff for every request.
func BackoffStrategy(strategy backoff.Strategy) PolicyOption {
	return func(pol *Policy) {
		pol.backoff = strategy
	}
}
