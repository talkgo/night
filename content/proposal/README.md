---
title: Proposing Changes to Go
---

## Introduction

The Go project's development process is design-driven.
Significant changes to the language, libraries, or tools must be first
discussed, and sometimes formally documented, before they can be implemented.

This document describes the process for proposing, documenting, and
implementing changes to the Go project.

To learn more about Go's origins and development process, see the talks
[How Go Was Made](https://talks.golang.org/2015/how-go-was-made.slide),
[The Evolution of Go](https://talks.golang.org/2015/gophercon-goevolution.slide),
and [Go, Open Source, Community](https://blog.golang.org/open-source)
from GopherCon 2015.

## The Proposal Process

The proposal process is the process for reviewing a proposal and reaching
a decision about whether to accept or decline the proposal.

1. The proposal author [creates a brief issue](https://golang.org/issue/new) describing the proposal.\
   Note: There is no need for a design document at this point.\
   Note: A non-proposal issue can be turned into a proposal by simply adding the proposal label.

2. A discussion on the issue tracker aims to triage the proposal into one of three outcomes:
     - Accept proposal, or
     - decline proposal, or
     - ask for a design doc.

   If the proposal is accepted or declined, the process is done.
   Otherwise the discussion is expected to identify concerns that
   should be addressed in a more detailed design.

3. The proposal author writes a [design doc](#design-documents) to work out details of the proposed
   design and address the concerns raised in the initial discussion.

4. Once comments and revisions on the design doc wind down, there is a final
   discussion on the issue, to reach one of two outcomes:
    - Accept proposal or
    - decline proposal.

After the proposal is accepted or declined (whether after step 2 or step 4),
implementation work proceeds in the same way as any other contribution.

## Detail

### Goals

- Make sure that proposals get a proper, fair, timely, recorded evaluation with
  a clear answer.
- Make past proposals easy to find, to avoid duplicated effort.
- If a design doc is needed, make sure contributors know how to write a good one.

### Definitions

- A **proposal** is a suggestion filed as a GitHub issue, identified by having
  the Proposal label.
- A **design doc** is the expanded form of a proposal, written when the
  proposal needs more careful explanation and consideration.

### Scope

The proposal process should be used for any notable change or addition to the
language, libraries and tools.
Since proposals begin (and will often end) with the filing of an issue, even
small changes can go through the proposal process if appropriate.
Deciding what is appropriate is matter of judgment we will refine through
experience.
If in doubt, file a proposal.

### Compatibility

Programs written for Go version 1.x must continue to compile and work with
future versions of Go 1.
The [Go 1 compatibility document](https://golang.org/doc/go1compat) describes
the promise we have made to Go users for the future of Go 1.x.
Any proposed change must not break this promise.

### Language changes

Go is a mature language and, as such, significant language changes are unlikely
to be accepted.
A "language change" in this context means a change to the
[Go language specification](https://golang.org/ref/spec).
(See the [release notes](https://golang.org/doc/devel/release.html) for
examples of recent language changes.)

### Design Documents

As noted above, some (but not all) proposals need to be elaborated in a design document.

- The design doc should be checked in to [the proposal repository](https://github.com/golang/proposal/) as `design/NNNN-shortname.md`,
where `NNNN` is the GitHub issue number and `shortname` is a short name
(a few dash-separated words at most).
Clone this repository with `git clone https://go.googlesource.com/proposal`
and follow the usual [Gerrit workflow for Go](https://golang.org/doc/contribute.html#Code_review).

- The design doc should follow [the template](design/TEMPLATE.md).

- The design doc should address any specific concerns raised during the initial discussion.

- It is expected that the design doc may go through multiple checked-in revisions.
New design doc authors may be paired with a design doc "shepherd" to help work on the doc.

- For ease of review with Gerrit, design documents should be wrapped around the
80 column mark.
[Each sentence should start on a new line](http://rhodesmill.org/brandon/2012/one-sentence-per-line/)
so that comments can be made accurately and the diff kept shorter.
  - In Emacs, loading `fill.el` from this directory will make `fill-paragraph` format text this way.

- Comments on Gerrit CLs should be restricted to grammar, spelling,
or procedural errors related to the preparation of the proposal itself.
All other comments should be addressed to the related GitHub issue.


### Quick Start for Experienced Committers

Experienced committers who are certain that a design doc will be
required for a particular proposal
can skip steps 1 and 2 and include the design doc with the initial issue.

In the worst case, skipping these steps only leads to an unnecessary design doc.

### Proposal Review

A group of Go team members holds “proposal review meetings”
approximately weekly to review pending proposals.

The principal goal of the review meeting is to make sure that proposals
are receiving attention from the right people,
by cc'ing relevant developers, raising important questions,
pinging lapsed discussions, and generally trying to guide discussion
toward agreement about the outcome.
The discussion itself is expected to happen on the issue tracker,
so that anyone can take part.

The proposal review meetings also identify issues where
consensus has been reached and the process can be
advanced to the next step (by marking the proposal accepted
or declined or by asking for a design doc).

### Consensus and Disagreement

The goal of the proposal process is to reach general consensus about the outcome
in a timely manner.

If general consensus cannot be reached,
the proposal review group decides the next step
by reviewing and discussing the issue and
reaching a consensus among themselves.
If even consensus among the proposal review group
cannot be reached (which would be exceedingly unusual),
the arbiter ([rsc@](mailto:rsc@golang.org))
reviews the discussion and
decides the next step.

## Help

If you need help with this process, please contact the Go contributors by posting
to the [golang-dev mailing list](https://groups.google.com/group/golang-dev).
(Note that the list is moderated, and that first-time posters should expect a
delay while their message is held for moderation.)

To learn about contributing to Go in general, see the
[contribution guidelines](https://golang.org/doc/contribute.html).