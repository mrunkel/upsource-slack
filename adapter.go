package main

import (
	"fmt"
	"strings"
)

var pathFeed = "%s/%s/review/%s?commentId=%s"
var pathReview = "%s/%s/review/%s"

//AdaptRequestResult adapt request to message
func AdaptRequestResult(request Request) Message {
	switch request.DataType {
	case "ReviewCreatedFeedEventBean":
		return buildReviewCreatedFeedEventBean(request)
	case "ReviewRemovedFeedEventBean":
		return buildReviewRemovedFeedEventBean(request)
	case "ReviewSquashedFeedEventBean":
		return buildReviewSquashedFeedEventBean(request)
	case "ReviewStateChangedFeedEventBean":
		return buildReviewStateChangedFeedEventBean(request)
	case "ReviewStoppedBranchTrackingFeedEventBean":
		return buildReviewStoppedBranchTrackingFeedEventBean(request)
	case "RevisionAddedToReviewFeedEventBean":
		return buildRevisionAddedToReviewFeedEventBean(request)
	case "RevisionRemovedFromReviewFeedEventBean":
		return buildRevisionRemovedFromReviewFeedEventBean(request)
	case "RemovedParticipantFromReviewFeedEventBean":
		return buildRemovedParticipantFromReviewFeedEventBean(request)
	case "PullRequestMergedFeedEventBean":
		return buildPullRequestMergedFeedEventBean(request)
	case "ParticipantStateChangedFeedEventBean":
		return buildParticipantStateChangedFeedEventBean(request)
	case "NewRevisionEventBean":
		return buildNewRevisionEventBean(request)
	case "NewParticipantInReviewFeedEventBean":
		return buildNewParticipantInReviewFeedEventBean(request)
	case "NewBranchEventBean":
		return buildNewBranchEventBean(request)
	case "MergedToDefaultBranchEventBean":
		return buildMergedToDefaultBranchEventBean(request)
	case "DiscussionFeedEventBean":
		return buildDiscussionFeedEventBean(request)
	}
	return Message{}
}
func buildReviewCreatedFeedEventBean(request Request) Message {
	message := Message{}
	data := request.Data.(ReviewCreatedFeedEventBean)
	attache := Attache{
		Author:    data.Base.Actor.UserName,
		Title:     "Review " + data.Base.ReviewID + " created",
		TitleLink: fmt.Sprintf(pathReview, request.Domain, request.ProjectID, data.Base.ReviewID),
		Text:      "Review " + data.Base.ReviewID + " created on branch" + data.Branch,
	}
	message.Attachments = append(message.Attachments, attache)
	return message
}

func buildReviewRemovedFeedEventBean(request Request) Message {
	message := Message{}
	data := request.Data.(ReviewRemovedFeedEventBean)
	attache := Attache{
		Author:    data.Base.Actor.UserName,
		Title:     "Review " + data.Base.ReviewID + " removed",
		TitleLink: fmt.Sprintf(pathReview, request.Domain, request.ProjectID, data.Base.ReviewID),
	}
	message.Attachments = append(message.Attachments, attache)
	return message
}

func buildReviewSquashedFeedEventBean(request Request) Message {
	message := Message{}
	data := request.Data.(ReviewSquashedFeedEventBean)
	attache := Attache{
		Author:    data.Base.Actor.UserName,
		Title:     "Review " + data.Base.ReviewID + " squashed ",
		TitleLink: fmt.Sprintf(pathReview, request.Domain, request.ProjectID, data.Base.ReviewID),
	}
	message.Attachments = append(message.Attachments, attache)
	return message
}

func buildReviewStateChangedFeedEventBean(request Request) Message {
	message := Message{}
	data := request.Data.(ReviewStateChangedFeedEventBean)
	attache := Attache{
		Author:    data.Base.Actor.UserName,
		Title:     "Review " + data.Base.ReviewID + " changed state ",
		TitleLink: fmt.Sprintf(pathReview, request.Domain, request.ProjectID, data.Base.ReviewID),
		Text:      data.Base.ReviewID + " review change state from " + string(data.OldState) + " to " + string(data.NewState),
	}
	message.Attachments = append(message.Attachments, attache)
	return message
}

func buildReviewStoppedBranchTrackingFeedEventBean(request Request) Message {
	message := Message{}
	data := request.Data.(ReviewStoppedBranchTrackingFeedEventBean)
	attache := Attache{
		Author:    data.Base.Actor.UserName,
		Title:     "Review stopped branch tracking " + data.Branch,
		TitleLink: fmt.Sprintf(pathReview, request.Domain, request.ProjectID, data.Base.ReviewID),
	}
	message.Attachments = append(message.Attachments, attache)
	return message
}
func buildRevisionAddedToReviewFeedEventBean(request Request) Message {
	message := Message{}
	data := request.Data.(RevisionAddedToReviewFeedEventBean)
	attache := Attache{
		Author:    data.Base.Actor.UserName,
		Title:     "Revisions " + strings.Join(data.RevisionIDs, ",") + " added from  " + data.Base.ReviewID,
		TitleLink: fmt.Sprintf(pathReview, request.Domain, request.ProjectID, data.Base.ReviewID),
	}
	message.Attachments = append(message.Attachments, attache)
	return message
}
func buildRevisionRemovedFromReviewFeedEventBean(request Request) Message {
	message := Message{}
	data := request.Data.(RevisionRemovedFromReviewFeedEventBean)
	attache := Attache{
		Author:    data.Base.Actor.UserName,
		Title:     "Revisions " + strings.Join(data.RevisionIDs, ",") + " removed from  " + data.Base.ReviewID,
		TitleLink: fmt.Sprintf(pathReview, request.Domain, request.ProjectID, data.Base.ReviewID),
	}
	message.Attachments = append(message.Attachments, attache)
	return message
}
func buildRemovedParticipantFromReviewFeedEventBean(request Request) Message {
	message := Message{}
	data := request.Data.(RemovedParticipantFromReviewFeedEventBean)
	attache := Attache{
		Author:    data.Base.Actor.UserName,
		Title:     data.Participant.UserName + " removed from  " + data.Base.ReviewID,
		TitleLink: fmt.Sprintf(pathReview, request.Domain, request.ProjectID, data.Base.ReviewID),
	}
	message.Attachments = append(message.Attachments, attache)
	return message
}

func buildPullRequestMergedFeedEventBean(request Request) Message {
	message := Message{}
	data := request.Data.(PullRequestMergedFeedEventBean)
	attache := Attache{
		Author:    data.Base.Actor.UserName,
		Title:     "Pull request merged  " + data.PullRequest,
		TitleLink: fmt.Sprintf(pathReview, request.Domain, request.ProjectID, data.Base.ReviewID),
	}
	message.Attachments = append(message.Attachments, attache)
	return message
}
func buildParticipantStateChangedFeedEventBean(request Request) Message {
	message := Message{}
	data := request.Data.(ParticipantStateChangedFeedEventBean)
	attache := Attache{
		Text:      data.Participant.UserName + "change state from " + string(data.OldState) + " to " + string(data.NewState),
		Author:    data.Participant.UserName,
		Title:     data.Participant.UserName + " change state",
		TitleLink: fmt.Sprintf(pathReview, request.Domain, request.ProjectID, data.Base.ReviewID),
	}
	message.Attachments = append(message.Attachments, attache)
	return message
}

func buildNewRevisionEventBean(request Request) Message {
	message := Message{}
	data := request.Data.(NewRevisionEventBean)
	attache := Attache{
		Text:   data.Message,
		Author: data.Author,
		Title:  "New revision  " + data.RevisionID,
	}
	message.Attachments = append(message.Attachments, attache)
	return message
}

func buildNewParticipantInReviewFeedEventBean(request Request) Message {
	message := Message{}
	data := request.Data.(NewParticipantInReviewFeedEventBean)
	attache := Attache{
		Text:      data.Participant.UserName,
		Title:     "New participant in review " + data.Base.ReviewID,
		TitleLink: fmt.Sprintf(pathReview, request.Domain, request.ProjectID, data.Base.ReviewID),
	}
	message.Attachments = append(message.Attachments, attache)
	return message
}

func buildNewBranchEventBean(request Request) Message {
	message := Message{}
	data := request.Data.(NewBranchEventBean)
	attache := Attache{
		Text:  data.Name,
		Title: "New branch",
	}
	message.Attachments = append(message.Attachments, attache)
	return message
}

func buildMergedToDefaultBranchEventBean(request Request) Message {
	message := Message{}
	data := request.Data.(MergedToDefaultBranchEventBean)
	attache := Attache{
		Text:  strings.Join(data.Branches, ","),
		Title: "Merged to default branch feed",
	}
	message.Attachments = append(message.Attachments, attache)
	return message
}

func buildDiscussionFeedEventBean(request Request) Message {
	message := Message{}
	data := request.Data.(DiscussionFeedEventBean)
	attache := Attache{
		Author:    data.Base.Actor.UserName,
		Text:      data.CommentText,
		Title:     "Discussion Feed " + data.Base.ReviewID,
		TitleLink: fmt.Sprintf(pathFeed, request.Domain, request.ProjectID, data.Base.ReviewID, data.CommentID),
		Color:     "#36a64f",
	}
	message.Attachments = append(message.Attachments, attache)
	return message
}
