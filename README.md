# id

The ID library is used to facilitate the creation of typed IDs that wrap a
common ID implementation and provide initialization functions that create
a instance of the typed ID.

The `id.ID` type is based on a google UUID, but hides the implementation
details from the library clients. 

`id.ID` implements the following methods which are made available to clients
that embed it:

- String() string to get the UUID string representation for the ID
- Scan(src any) to support reading IDs from a databse
- Value() (driver.Value, error) to support reading IDs from a database
- MarshalJSON() ([]byte, error) to serialize an ID as JSON
- UnmarshalJSON([]byte) error to parse a JSONstring as an ID
- IsNil() bool to determine if the ID is nil/unset (all zeros)

For example, if a new ID type called `EventID` was required, an `EventID`
struct would be created that embeds an `id.ID` and the `id.Initializers()`
function would be invoked to create `EventID` initializer functions as follows:

```
import "github.com/dmpettyp/id"

// Define the EventID type, and embed an id.ID in it
type EventID struct{ id.ID }

// Create the initializers for the EventID type. To create the initializers
// a function must be provided that returns the new ID type with the embedded
// id.ID set in it.
var NewEventID, MustNewEventID, ParseEventID = id.Inititalizers(
	func(id id.ID) EventID { return EventID{ID: id} },
)

// NewEventID has the signature NewEventID() (EventID, error) and will return
// an error if a new EventID cannot initialized

// MustNewEventID has the signature NewEventID() EventID and will panic if a 
// new EventID cannot initialized

// Parse has the signature NewEventID() EventID and will panic if a 
// new EventID cannot initialized
```
