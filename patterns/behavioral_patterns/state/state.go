package main

import "fmt"

type State interface {
	review(document *Document)
	publish(document *Document)
}

type Document struct {
	state State
}

func (d *Document) review() {
	d.state.review(d)
}

func (d *Document) publish() {
	d.state.publish(d)
}

func (d *Document) setState(state State) {
	d.state = state
}

func (d *Document) getCurrentState() State {
	return d.state
}

type DraftState struct{}

func (s *DraftState) review(document *Document) {
	fmt.Println("Document reviewed, moving to moderation.")
	document.setState(&ModerationState{})
}

func (s *DraftState) publish(document *Document) {
	fmt.Println("Draft document cannot be published directly.")
}

type ModerationState struct{}

func (s *ModerationState) review(document *Document) {
	fmt.Println("Document is already under review.")
}

func (s *ModerationState) publish(document *Document) {
	fmt.Println("Document published successfully.")
	document.setState(&PublishedState{})
}

type PublishedState struct{}

func (s *PublishedState) review(document *Document) {
	fmt.Println("Published document cannot be reviewed again.")
}

func (s *PublishedState) publish(document *Document) {
	fmt.Println("Document is already published.")
}

func main() {
	document := &Document{state: &DraftState{}}

	document.review()
	document.publish()
	document.review()
	document.publish()
}
