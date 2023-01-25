package api

import (
	"encoding/json"

	"github.com/bytebase/bytebase/backend/common"
)

// TaskRunStatus is the status of a task run.
type TaskRunStatus string

const (
	// TaskRunUnknown is the task run status for UNKNOWN.
	TaskRunUnknown TaskRunStatus = "UNKNOWN"
	// TaskRunRunning is the task run status for RUNNING.
	TaskRunRunning TaskRunStatus = "RUNNING"
	// TaskRunDone is the task run status for DONE.
	TaskRunDone TaskRunStatus = "DONE"
	// TaskRunFailed is the task run status for FAILED.
	TaskRunFailed TaskRunStatus = "FAILED"
	// TaskRunCanceled is the task run status for CANCELED.
	TaskRunCanceled TaskRunStatus = "CANCELED"
)

// TaskRunResultPayload is the result payload for a task run.
type TaskRunResultPayload struct {
	Detail      string `json:"detail,omitempty"`
	MigrationID string `json:"migrationId,omitempty"`
	Version     string `json:"version,omitempty"`
}

// TaskRun is the API message for a task run.
type TaskRun struct {
	ID int `jsonapi:"primary,taskRun"`

	// Standard fields
	CreatorID int
	Creator   *Principal `jsonapi:"relation,creator"`
	CreatedTs int64      `jsonapi:"attr,createdTs"`
	UpdaterID int
	Updater   *Principal `jsonapi:"relation,updater"`
	UpdatedTs int64      `jsonapi:"attr,updatedTs"`

	// Related fields
	TaskID int `jsonapi:"attr,taskId"`

	// Domain specific fields
	Name    string        `jsonapi:"attr,name"`
	Status  TaskRunStatus `jsonapi:"attr,status"`
	Type    TaskType      `jsonapi:"attr,type"`
	Code    common.Code   `jsonapi:"attr,code"`
	Comment string        `jsonapi:"attr,comment"`
	Result  string        `jsonapi:"attr,result"`
	Payload string        `jsonapi:"attr,payload"`
}

// TaskRunCreate is the API message for creating a task run.
type TaskRunCreate struct {
	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	CreatorID int

	// Related fields
	TaskID int

	// Domain specific fields
	Name    string   `jsonapi:"attr,name"`
	Type    TaskType `jsonapi:"attr,type"`
	Payload string   `jsonapi:"attr,payload"`
}

// TaskRunFind is the API message for finding task runs.
type TaskRunFind struct {
	ID *int

	// Related fields
	TaskID *int

	// Domain specific fields
	StatusList *[]TaskRunStatus
}

func (find *TaskRunFind) String() string {
	str, err := json.Marshal(*find)
	if err != nil {
		return err.Error()
	}
	return string(str)
}

// TaskRunStatusPatch is the API message for patching a task run.
type TaskRunStatusPatch struct {
	ID *int

	// Standard fields
	UpdaterID int

	// Related fields
	TaskID *int

	// Domain specific fields
	Status TaskRunStatus
	Code   *common.Code
	// Records the status detail (e.g. error message on failure)
	Comment *string
	Result  *string
}