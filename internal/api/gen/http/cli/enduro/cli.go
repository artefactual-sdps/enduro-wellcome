// Code generated by goa v3.8.1, DO NOT EDIT.
//
// enduro HTTP client CLI support package
//
// Command:
// $ goa-v3.8.1 gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	batchc "github.com/artefactual-sdps/enduro/internal/api/gen/http/batch/client"
	package_c "github.com/artefactual-sdps/enduro/internal/api/gen/http/package_/client"
	storagec "github.com/artefactual-sdps/enduro/internal/api/gen/http/storage/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `batch (submit|status|hints)
package (monitor|list|show|delete|cancel|retry|bulk|bulk-status|preservation-actions|confirm|reject|move|move-status)
storage (submit|update|download|list|move|move-status|reject|show)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` batch submit --body '{
      "completed_dir": "Vitae odit sunt.",
      "path": "Illum quasi.",
      "retention_period": "Dolor et suscipit."
   }'` + "\n" +
		os.Args[0] + ` package monitor` + "\n" +
		os.Args[0] + ` storage submit --body '{
      "name": "Explicabo qui."
   }' --aip-id "Omnis quod officiis rem voluptas."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
	dialer goahttp.Dialer,
	package_Configurer *package_c.ConnConfigurer,
) (goa.Endpoint, interface{}, error) {
	var (
		batchFlags = flag.NewFlagSet("batch", flag.ContinueOnError)

		batchSubmitFlags    = flag.NewFlagSet("submit", flag.ExitOnError)
		batchSubmitBodyFlag = batchSubmitFlags.String("body", "REQUIRED", "")

		batchStatusFlags = flag.NewFlagSet("status", flag.ExitOnError)

		batchHintsFlags = flag.NewFlagSet("hints", flag.ExitOnError)

		package_Flags = flag.NewFlagSet("package", flag.ContinueOnError)

		package_MonitorFlags = flag.NewFlagSet("monitor", flag.ExitOnError)

		package_ListFlags                   = flag.NewFlagSet("list", flag.ExitOnError)
		package_ListNameFlag                = package_ListFlags.String("name", "", "")
		package_ListAipIDFlag               = package_ListFlags.String("aip-id", "", "")
		package_ListEarliestCreatedTimeFlag = package_ListFlags.String("earliest-created-time", "", "")
		package_ListLatestCreatedTimeFlag   = package_ListFlags.String("latest-created-time", "", "")
		package_ListLocationFlag            = package_ListFlags.String("location", "", "")
		package_ListStatusFlag              = package_ListFlags.String("status", "", "")
		package_ListCursorFlag              = package_ListFlags.String("cursor", "", "")

		package_ShowFlags  = flag.NewFlagSet("show", flag.ExitOnError)
		package_ShowIDFlag = package_ShowFlags.String("id", "REQUIRED", "Identifier of package to show")

		package_DeleteFlags  = flag.NewFlagSet("delete", flag.ExitOnError)
		package_DeleteIDFlag = package_DeleteFlags.String("id", "REQUIRED", "Identifier of package to delete")

		package_CancelFlags  = flag.NewFlagSet("cancel", flag.ExitOnError)
		package_CancelIDFlag = package_CancelFlags.String("id", "REQUIRED", "Identifier of package to remove")

		package_RetryFlags  = flag.NewFlagSet("retry", flag.ExitOnError)
		package_RetryIDFlag = package_RetryFlags.String("id", "REQUIRED", "Identifier of package to retry")

		package_BulkFlags    = flag.NewFlagSet("bulk", flag.ExitOnError)
		package_BulkBodyFlag = package_BulkFlags.String("body", "REQUIRED", "")

		package_BulkStatusFlags = flag.NewFlagSet("bulk-status", flag.ExitOnError)

		package_PreservationActionsFlags  = flag.NewFlagSet("preservation-actions", flag.ExitOnError)
		package_PreservationActionsIDFlag = package_PreservationActionsFlags.String("id", "REQUIRED", "Identifier of package to look up")

		package_ConfirmFlags    = flag.NewFlagSet("confirm", flag.ExitOnError)
		package_ConfirmBodyFlag = package_ConfirmFlags.String("body", "REQUIRED", "")
		package_ConfirmIDFlag   = package_ConfirmFlags.String("id", "REQUIRED", "Identifier of package to look up")

		package_RejectFlags  = flag.NewFlagSet("reject", flag.ExitOnError)
		package_RejectIDFlag = package_RejectFlags.String("id", "REQUIRED", "Identifier of package to look up")

		package_MoveFlags    = flag.NewFlagSet("move", flag.ExitOnError)
		package_MoveBodyFlag = package_MoveFlags.String("body", "REQUIRED", "")
		package_MoveIDFlag   = package_MoveFlags.String("id", "REQUIRED", "Identifier of package to move")

		package_MoveStatusFlags  = flag.NewFlagSet("move-status", flag.ExitOnError)
		package_MoveStatusIDFlag = package_MoveStatusFlags.String("id", "REQUIRED", "Identifier of package to move")

		storageFlags = flag.NewFlagSet("storage", flag.ContinueOnError)

		storageSubmitFlags     = flag.NewFlagSet("submit", flag.ExitOnError)
		storageSubmitBodyFlag  = storageSubmitFlags.String("body", "REQUIRED", "")
		storageSubmitAipIDFlag = storageSubmitFlags.String("aip-id", "REQUIRED", "")

		storageUpdateFlags     = flag.NewFlagSet("update", flag.ExitOnError)
		storageUpdateAipIDFlag = storageUpdateFlags.String("aip-id", "REQUIRED", "")

		storageDownloadFlags     = flag.NewFlagSet("download", flag.ExitOnError)
		storageDownloadAipIDFlag = storageDownloadFlags.String("aip-id", "REQUIRED", "")

		storageListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		storageMoveFlags     = flag.NewFlagSet("move", flag.ExitOnError)
		storageMoveBodyFlag  = storageMoveFlags.String("body", "REQUIRED", "")
		storageMoveAipIDFlag = storageMoveFlags.String("aip-id", "REQUIRED", "")

		storageMoveStatusFlags     = flag.NewFlagSet("move-status", flag.ExitOnError)
		storageMoveStatusAipIDFlag = storageMoveStatusFlags.String("aip-id", "REQUIRED", "")

		storageRejectFlags     = flag.NewFlagSet("reject", flag.ExitOnError)
		storageRejectAipIDFlag = storageRejectFlags.String("aip-id", "REQUIRED", "")

		storageShowFlags     = flag.NewFlagSet("show", flag.ExitOnError)
		storageShowAipIDFlag = storageShowFlags.String("aip-id", "REQUIRED", "")
	)
	batchFlags.Usage = batchUsage
	batchSubmitFlags.Usage = batchSubmitUsage
	batchStatusFlags.Usage = batchStatusUsage
	batchHintsFlags.Usage = batchHintsUsage

	package_Flags.Usage = package_Usage
	package_MonitorFlags.Usage = package_MonitorUsage
	package_ListFlags.Usage = package_ListUsage
	package_ShowFlags.Usage = package_ShowUsage
	package_DeleteFlags.Usage = package_DeleteUsage
	package_CancelFlags.Usage = package_CancelUsage
	package_RetryFlags.Usage = package_RetryUsage
	package_BulkFlags.Usage = package_BulkUsage
	package_BulkStatusFlags.Usage = package_BulkStatusUsage
	package_PreservationActionsFlags.Usage = package_PreservationActionsUsage
	package_ConfirmFlags.Usage = package_ConfirmUsage
	package_RejectFlags.Usage = package_RejectUsage
	package_MoveFlags.Usage = package_MoveUsage
	package_MoveStatusFlags.Usage = package_MoveStatusUsage

	storageFlags.Usage = storageUsage
	storageSubmitFlags.Usage = storageSubmitUsage
	storageUpdateFlags.Usage = storageUpdateUsage
	storageDownloadFlags.Usage = storageDownloadUsage
	storageListFlags.Usage = storageListUsage
	storageMoveFlags.Usage = storageMoveUsage
	storageMoveStatusFlags.Usage = storageMoveStatusUsage
	storageRejectFlags.Usage = storageRejectUsage
	storageShowFlags.Usage = storageShowUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "batch":
			svcf = batchFlags
		case "package":
			svcf = package_Flags
		case "storage":
			svcf = storageFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "batch":
			switch epn {
			case "submit":
				epf = batchSubmitFlags

			case "status":
				epf = batchStatusFlags

			case "hints":
				epf = batchHintsFlags

			}

		case "package":
			switch epn {
			case "monitor":
				epf = package_MonitorFlags

			case "list":
				epf = package_ListFlags

			case "show":
				epf = package_ShowFlags

			case "delete":
				epf = package_DeleteFlags

			case "cancel":
				epf = package_CancelFlags

			case "retry":
				epf = package_RetryFlags

			case "bulk":
				epf = package_BulkFlags

			case "bulk-status":
				epf = package_BulkStatusFlags

			case "preservation-actions":
				epf = package_PreservationActionsFlags

			case "confirm":
				epf = package_ConfirmFlags

			case "reject":
				epf = package_RejectFlags

			case "move":
				epf = package_MoveFlags

			case "move-status":
				epf = package_MoveStatusFlags

			}

		case "storage":
			switch epn {
			case "submit":
				epf = storageSubmitFlags

			case "update":
				epf = storageUpdateFlags

			case "download":
				epf = storageDownloadFlags

			case "list":
				epf = storageListFlags

			case "move":
				epf = storageMoveFlags

			case "move-status":
				epf = storageMoveStatusFlags

			case "reject":
				epf = storageRejectFlags

			case "show":
				epf = storageShowFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "batch":
			c := batchc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "submit":
				endpoint = c.Submit()
				data, err = batchc.BuildSubmitPayload(*batchSubmitBodyFlag)
			case "status":
				endpoint = c.Status()
				data = nil
			case "hints":
				endpoint = c.Hints()
				data = nil
			}
		case "package":
			c := package_c.NewClient(scheme, host, doer, enc, dec, restore, dialer, package_Configurer)
			switch epn {
			case "monitor":
				endpoint = c.Monitor()
				data = nil
			case "list":
				endpoint = c.List()
				data, err = package_c.BuildListPayload(*package_ListNameFlag, *package_ListAipIDFlag, *package_ListEarliestCreatedTimeFlag, *package_ListLatestCreatedTimeFlag, *package_ListLocationFlag, *package_ListStatusFlag, *package_ListCursorFlag)
			case "show":
				endpoint = c.Show()
				data, err = package_c.BuildShowPayload(*package_ShowIDFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = package_c.BuildDeletePayload(*package_DeleteIDFlag)
			case "cancel":
				endpoint = c.Cancel()
				data, err = package_c.BuildCancelPayload(*package_CancelIDFlag)
			case "retry":
				endpoint = c.Retry()
				data, err = package_c.BuildRetryPayload(*package_RetryIDFlag)
			case "bulk":
				endpoint = c.Bulk()
				data, err = package_c.BuildBulkPayload(*package_BulkBodyFlag)
			case "bulk-status":
				endpoint = c.BulkStatus()
				data = nil
			case "preservation-actions":
				endpoint = c.PreservationActions()
				data, err = package_c.BuildPreservationActionsPayload(*package_PreservationActionsIDFlag)
			case "confirm":
				endpoint = c.Confirm()
				data, err = package_c.BuildConfirmPayload(*package_ConfirmBodyFlag, *package_ConfirmIDFlag)
			case "reject":
				endpoint = c.Reject()
				data, err = package_c.BuildRejectPayload(*package_RejectIDFlag)
			case "move":
				endpoint = c.Move()
				data, err = package_c.BuildMovePayload(*package_MoveBodyFlag, *package_MoveIDFlag)
			case "move-status":
				endpoint = c.MoveStatus()
				data, err = package_c.BuildMoveStatusPayload(*package_MoveStatusIDFlag)
			}
		case "storage":
			c := storagec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "submit":
				endpoint = c.Submit()
				data, err = storagec.BuildSubmitPayload(*storageSubmitBodyFlag, *storageSubmitAipIDFlag)
			case "update":
				endpoint = c.Update()
				data, err = storagec.BuildUpdatePayload(*storageUpdateAipIDFlag)
			case "download":
				endpoint = c.Download()
				data, err = storagec.BuildDownloadPayload(*storageDownloadAipIDFlag)
			case "list":
				endpoint = c.List()
				data = nil
			case "move":
				endpoint = c.Move()
				data, err = storagec.BuildMovePayload(*storageMoveBodyFlag, *storageMoveAipIDFlag)
			case "move-status":
				endpoint = c.MoveStatus()
				data, err = storagec.BuildMoveStatusPayload(*storageMoveStatusAipIDFlag)
			case "reject":
				endpoint = c.Reject()
				data, err = storagec.BuildRejectPayload(*storageRejectAipIDFlag)
			case "show":
				endpoint = c.Show()
				data, err = storagec.BuildShowPayload(*storageShowAipIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// batchUsage displays the usage of the batch command and its subcommands.
func batchUsage() {
	fmt.Fprintf(os.Stderr, `The batch service manages batches of packages.
Usage:
    %[1]s [globalflags] batch COMMAND [flags]

COMMAND:
    submit: Submit a new batch
    status: Retrieve status of current batch operation.
    hints: Retrieve form hints

Additional help:
    %[1]s batch COMMAND --help
`, os.Args[0])
}
func batchSubmitUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] batch submit -body JSON

Submit a new batch
    -body JSON: 

Example:
    %[1]s batch submit --body '{
      "completed_dir": "Vitae odit sunt.",
      "path": "Illum quasi.",
      "retention_period": "Dolor et suscipit."
   }'
`, os.Args[0])
}

func batchStatusUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] batch status

Retrieve status of current batch operation.

Example:
    %[1]s batch status
`, os.Args[0])
}

func batchHintsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] batch hints

Retrieve form hints

Example:
    %[1]s batch hints
`, os.Args[0])
}

// packageUsage displays the usage of the package command and its subcommands.
func package_Usage() {
	fmt.Fprintf(os.Stderr, `The package service manages packages being transferred to a3m.
Usage:
    %[1]s [globalflags] package COMMAND [flags]

COMMAND:
    monitor: Monitor implements monitor.
    list: List all stored packages
    show: Show package by ID
    delete: Delete package by ID
    cancel: Cancel package processing by ID
    retry: Retry package processing by ID
    bulk: Bulk operations (retry, cancel...).
    bulk-status: Retrieve status of current bulk operation.
    preservation-actions: List all preservation actions by ID
    confirm: Signal the package has been reviewed and accepted
    reject: Signal the package has been reviewed and rejected
    move: Move a package to a permanent storage location
    move-status: Retrieve the status of a permanent storage location move of the package

Additional help:
    %[1]s package COMMAND --help
`, os.Args[0])
}
func package_MonitorUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package monitor

Monitor implements monitor.

Example:
    %[1]s package monitor
`, os.Args[0])
}

func package_ListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package list -name STRING -aip-id STRING -earliest-created-time STRING -latest-created-time STRING -location STRING -status STRING -cursor STRING

List all stored packages
    -name STRING: 
    -aip-id STRING: 
    -earliest-created-time STRING: 
    -latest-created-time STRING: 
    -location STRING: 
    -status STRING: 
    -cursor STRING: 

Example:
    %[1]s package list --name "Debitis ab aliquid est reprehenderit." --aip-id "A8E7C55C-B3EA-B228-EF84-9DB3D55D57A4" --earliest-created-time "1989-04-22T06:53:55Z" --latest-created-time "2014-12-22T11:44:55Z" --location "Ea in quia atque commodi et." --status "error" --cursor "Minus velit magnam recusandae modi dignissimos."
`, os.Args[0])
}

func package_ShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package show -id UINT

Show package by ID
    -id UINT: Identifier of package to show

Example:
    %[1]s package show --id 9340392832192809768
`, os.Args[0])
}

func package_DeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package delete -id UINT

Delete package by ID
    -id UINT: Identifier of package to delete

Example:
    %[1]s package delete --id 10916720185593198763
`, os.Args[0])
}

func package_CancelUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package cancel -id UINT

Cancel package processing by ID
    -id UINT: Identifier of package to remove

Example:
    %[1]s package cancel --id 646974705416522731
`, os.Args[0])
}

func package_RetryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package retry -id UINT

Retry package processing by ID
    -id UINT: Identifier of package to retry

Example:
    %[1]s package retry --id 2018151398983385474
`, os.Args[0])
}

func package_BulkUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package bulk -body JSON

Bulk operations (retry, cancel...).
    -body JSON: 

Example:
    %[1]s package bulk --body '{
      "operation": "abandon",
      "size": 17552965550942373114,
      "status": "unknown"
   }'
`, os.Args[0])
}

func package_BulkStatusUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package bulk-status

Retrieve status of current bulk operation.

Example:
    %[1]s package bulk-status
`, os.Args[0])
}

func package_PreservationActionsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package preservation-actions -id UINT

List all preservation actions by ID
    -id UINT: Identifier of package to look up

Example:
    %[1]s package preservation-actions --id 7612789260457193192
`, os.Args[0])
}

func package_ConfirmUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package confirm -body JSON -id UINT

Signal the package has been reviewed and accepted
    -body JSON: 
    -id UINT: Identifier of package to look up

Example:
    %[1]s package confirm --body '{
      "location": "Ullam eos eius officiis rerum assumenda."
   }' --id 14869185368804314752
`, os.Args[0])
}

func package_RejectUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package reject -id UINT

Signal the package has been reviewed and rejected
    -id UINT: Identifier of package to look up

Example:
    %[1]s package reject --id 6758623569740053693
`, os.Args[0])
}

func package_MoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package move -body JSON -id UINT

Move a package to a permanent storage location
    -body JSON: 
    -id UINT: Identifier of package to move

Example:
    %[1]s package move --body '{
      "location": "Officia quibusdam dolore in aliquid aut optio."
   }' --id 3875033090007834170
`, os.Args[0])
}

func package_MoveStatusUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package move-status -id UINT

Retrieve the status of a permanent storage location move of the package
    -id UINT: Identifier of package to move

Example:
    %[1]s package move-status --id 16681095185899787694
`, os.Args[0])
}

// storageUsage displays the usage of the storage command and its subcommands.
func storageUsage() {
	fmt.Fprintf(os.Stderr, `The storage service manages the storage of packages.
Usage:
    %[1]s [globalflags] storage COMMAND [flags]

COMMAND:
    submit: Start the submission of a package
    update: Signal the storage service that an upload is complete
    download: Download package by AIPID
    list: List locations
    move: Move a package to a permanent storage location
    move-status: Retrieve the status of a permanent storage location move of the package
    reject: Reject a package
    show: Show package by AIPID

Additional help:
    %[1]s storage COMMAND --help
`, os.Args[0])
}
func storageSubmitUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage submit -body JSON -aip-id STRING

Start the submission of a package
    -body JSON: 
    -aip-id STRING: 

Example:
    %[1]s storage submit --body '{
      "name": "Explicabo qui."
   }' --aip-id "Omnis quod officiis rem voluptas."
`, os.Args[0])
}

func storageUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage update -aip-id STRING

Signal the storage service that an upload is complete
    -aip-id STRING: 

Example:
    %[1]s storage update --aip-id "Porro libero consequatur commodi reprehenderit."
`, os.Args[0])
}

func storageDownloadUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage download -aip-id STRING

Download package by AIPID
    -aip-id STRING: 

Example:
    %[1]s storage download --aip-id "Porro numquam dolores doloribus."
`, os.Args[0])
}

func storageListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage list

List locations

Example:
    %[1]s storage list
`, os.Args[0])
}

func storageMoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage move -body JSON -aip-id STRING

Move a package to a permanent storage location
    -body JSON: 
    -aip-id STRING: 

Example:
    %[1]s storage move --body '{
      "location": "Eaque architecto magnam pariatur rerum voluptas."
   }' --aip-id "Sit vitae."
`, os.Args[0])
}

func storageMoveStatusUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage move-status -aip-id STRING

Retrieve the status of a permanent storage location move of the package
    -aip-id STRING: 

Example:
    %[1]s storage move-status --aip-id "Dolorum quis sed iure mollitia nisi."
`, os.Args[0])
}

func storageRejectUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage reject -aip-id STRING

Reject a package
    -aip-id STRING: 

Example:
    %[1]s storage reject --aip-id "Mollitia repellendus et qui ratione esse."
`, os.Args[0])
}

func storageShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage show -aip-id STRING

Show package by AIPID
    -aip-id STRING: 

Example:
    %[1]s storage show --aip-id "Unde saepe."
`, os.Args[0])
}
