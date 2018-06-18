package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

// Request request
type Request struct {
	Domain       string
	MajorVersion int
	MinorVersion int
	ProjectID    string
	DataType     string
	Data         interface{}
}

// ReviewCreatedFeedEventBean request
type ReviewCreatedFeedEventBean struct {
	Base      FeedEventBean
	Revisions []string
	Branch    string
}

// ReviewRemovedFeedEventBean request
type ReviewRemovedFeedEventBean struct {
	Base     FeedEventBean
	ReviewID string
}

// ReviewSquashedFeedEventBean request
type ReviewSquashedFeedEventBean struct {
	Base       FeedEventBean
	RevisionID string
}

// ReviewStateChangedFeedEventBean request
type ReviewStateChangedFeedEventBean struct {
	Base     FeedEventBean
	OldState ReviewState
	NewState ReviewState
}

// ReviewStoppedBranchTrackingFeedEventBean request
type ReviewStoppedBranchTrackingFeedEventBean struct {
	Base   FeedEventBean
	Branch string
}

// RevisionAddedToReviewFeedEventBean request
type RevisionAddedToReviewFeedEventBean struct {
	Base        FeedEventBean
	RevisionID  string
	RevisionIDs []string
}

// RevisionRemovedFromReviewFeedEventBean request
type RevisionRemovedFromReviewFeedEventBean struct {
	Base        FeedEventBean
	RevisionIDs []string
}

// UserIDBean request
type UserIDBean struct {
	UserID    string
	UserName  string
	UserEmail string
}

// RemovedParticipantFromReviewFeedEventBean request
type RemovedParticipantFromReviewFeedEventBean struct {
	Base        FeedEventBean
	Participant UserIDBean
	FormerRole  ParticipantRole
}

// PullRequestMergedFeedEventBean request
type PullRequestMergedFeedEventBean struct {
	Base        FeedEventBean
	PullRequest string
}

// ParticipantStateChangedFeedEventBean request
type ParticipantStateChangedFeedEventBean struct {
	Base        FeedEventBean
	Participant UserIDBean
	OldState    ParticipantState
	NewState    ParticipantState
}

// NewRevisionEventBean request
type NewRevisionEventBean struct {
	RevisionID string
	Branches   []string
	Author     string
	Message    string
	Date       int64
}

// NewParticipantInReviewFeedEventBean request
type NewParticipantInReviewFeedEventBean struct {
	Base        FeedEventBean
	Participant UserIDBean
	Role        ParticipantRole
}

// NewBranchEventBean request
type NewBranchEventBean struct {
	Name string
}

// MergedToDefaultBranchEventBean request
type MergedToDefaultBranchEventBean struct {
	CommitsCount int32
	Branches     []string
}

// FeedEventBean request
type FeedEventBean struct {
	UserID       UserIDBean
	UserIDs      []UserIDBean
	ReviewNumber int32
	ReviewID     string
	Date         int64
	Actor        UserIDBean
	FeedEventID  string
}

// DiscussionFeedEventBean request
type DiscussionFeedEventBean struct {
	Base               FeedEventBean
	NotificationReason NotificationReason
	DiscussionID       string
	CommentID          string
	CommentText        string
	IsEdit             bool
	ResolveAction      bool
}

// NotificationReason request
type NotificationReason string

// ParticipantRole request
type ParticipantRole string

// ParticipantState request
type ParticipantState string

// ReviewState request
type ReviewState string

var typeMap = map[string]reflect.Type{
	"ReviewCreatedFeedEventBean":                reflect.TypeOf(ReviewCreatedFeedEventBean{}),
	"ReviewRemovedFeedEventBean":                reflect.TypeOf(ReviewRemovedFeedEventBean{}),
	"ReviewSquashedFeedEventBean":               reflect.TypeOf(ReviewSquashedFeedEventBean{}),
	"ReviewStateChangedFeedEventBean":           reflect.TypeOf(ReviewStateChangedFeedEventBean{}),
	"ReviewStoppedBranchTrackingFeedEventBean":  reflect.TypeOf(ReviewStoppedBranchTrackingFeedEventBean{}),
	"RevisionAddedToReviewFeedEventBean":        reflect.TypeOf(RevisionAddedToReviewFeedEventBean{}),
	"RevisionRemovedFromReviewFeedEventBean":    reflect.TypeOf(RevisionRemovedFromReviewFeedEventBean{}),
	"RemovedParticipantFromReviewFeedEventBean": reflect.TypeOf(RemovedParticipantFromReviewFeedEventBean{}),
	"PullRequestMergedFeedEventBean":            reflect.TypeOf(PullRequestMergedFeedEventBean{}),
	"ParticipantStateChangedFeedEventBean":      reflect.TypeOf(ParticipantStateChangedFeedEventBean{}),
	"NewRevisionEventBean":                      reflect.TypeOf(NewRevisionEventBean{}),
	"NewParticipantInReviewFeedEventBean":       reflect.TypeOf(NewParticipantInReviewFeedEventBean{}),
	"NewBranchEventBean":                        reflect.TypeOf(NewBranchEventBean{}),
	"MergedToDefaultBranchEventBean":            reflect.TypeOf(MergedToDefaultBranchEventBean{}),
	"DiscussionFeedEventBean":                   reflect.TypeOf(DiscussionFeedEventBean{}),
}

//Parse parse request upsource
func Parse(body []byte, domain string) Request {
	var request Request
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Panic(err)
	}
	result := getResult(request)
	request.Data = result
	request.Domain = domain
	return request
}

func getResult(request Request) interface{} {
	switch request.DataType {
	case "ReviewCreatedFeedEventBean":
		var result NewRevisionEventBean
		byteData, _ := json.Marshal(request.Data)
		json.Unmarshal(byteData, &result)
		return result
	case "ReviewRemovedFeedEventBean":
		var result ReviewRemovedFeedEventBean
		byteData, _ := json.Marshal(request.Data)
		json.Unmarshal(byteData, &result)
		return result
	case "ReviewSquashedFeedEventBean":
		var result ReviewSquashedFeedEventBean
		byteData, _ := json.Marshal(request.Data)
		json.Unmarshal(byteData, &result)
		return result
	case "ReviewStateChangedFeedEventBean":
		var result ReviewStateChangedFeedEventBean
		byteData, _ := json.Marshal(request.Data)
		json.Unmarshal(byteData, &result)
		return result
	case "ReviewStoppedBranchTrackingFeedEventBean":
		var result ReviewStoppedBranchTrackingFeedEventBean
		byteData, _ := json.Marshal(request.Data)
		json.Unmarshal(byteData, &result)
		return result
	case "RevisionAddedToReviewFeedEventBean":
		var result RevisionAddedToReviewFeedEventBean
		byteData, _ := json.Marshal(request.Data)
		json.Unmarshal(byteData, &result)
		return result
	case "RevisionRemovedFromReviewFeedEventBean":
		var result RevisionRemovedFromReviewFeedEventBean
		byteData, _ := json.Marshal(request.Data)
		json.Unmarshal(byteData, &result)
		return result
	case "RemovedParticipantFromReviewFeedEventBean":
		var result RemovedParticipantFromReviewFeedEventBean
		byteData, _ := json.Marshal(request.Data)
		json.Unmarshal(byteData, &result)
		return result
	case "PullRequestMergedFeedEventBean":
		var result PullRequestMergedFeedEventBean
		byteData, _ := json.Marshal(request.Data)
		json.Unmarshal(byteData, &result)
		return result
	case "ParticipantStateChangedFeedEventBean":
		var result ParticipantStateChangedFeedEventBean
		byteData, _ := json.Marshal(request.Data)
		json.Unmarshal(byteData, &result)
		return result
	case "NewRevisionEventBean":
		var result NewRevisionEventBean
		byteData, _ := json.Marshal(request.Data)
		json.Unmarshal(byteData, &result)
		return result
	case "NewParticipantInReviewFeedEventBean":
		var result NewParticipantInReviewFeedEventBean
		byteData, _ := json.Marshal(request.Data)
		json.Unmarshal(byteData, &result)
		return result
	case "NewBranchEventBean":
		var result NewBranchEventBean
		byteData, _ := json.Marshal(request.Data)
		json.Unmarshal(byteData, &result)
		return result
	case "MergedToDefaultBranchEventBean":
		var result MergedToDefaultBranchEventBean
		byteData, _ := json.Marshal(request.Data)
		json.Unmarshal(byteData, &result)
		return result
	case "DiscussionFeedEventBean":
		var result DiscussionFeedEventBean
		byteData, _ := json.Marshal(request.Data)
		json.Unmarshal(byteData, &result)
		return result
	}
	return nil
}

func (a *NotificationReason) UnmarshalJSON(b []byte) error {
	var s int
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	fmt.Println(s)
	switch s {
	case 0:
		*a = "Unknown"
	case 1:
		*a = "CommentInAuthorFeed"
	case 2:
		*a = "NotifyCommitAuthor"
	case 3:
		*a = "Mention"
	case 4:
		*a = "HashTagSubscription"
	case 5:
		*a = "DiscussionIsStarred"
	case 6:
		*a = "ParticipatedInDiscussion"
	case 7:
		*a = "ParticipatedInReview"
	case 8:
		*a = "CommentInAuthorFeed"
	}

	return nil
}

func (a *ParticipantRole) UnmarshalJSON(b []byte) error {
	var s int
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case 1:
		*a = "Author"
	case 2:
		*a = "Reviewer"
	case 3:
		*a = "Watcher"
	}

	return nil
}

func (a *ParticipantState) UnmarshalJSON(b []byte) error {
	var s int
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case 0:
		*a = "Unread"
	case 1:
		*a = "Read"
	case 2:
		*a = "Accepted"
	case 3:
		*a = "Rejected"
	}

	return nil
}

func (a *ReviewState) UnmarshalJSON(b []byte) error {
	var s int
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case 0:
		*a = "Open"
	case 1:
		*a = "Closed"
	}

	return nil
}
