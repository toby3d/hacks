package http

// This status code is primary meant to be used in REST API interactions as an
// issued warning related to the request itself. Since in regular GET requests
// there are no direct ways of telling the client about a potential mistake in
// the request without altering the response body, the only way to notify the
// client is using a custom HTTP header.
//
// However, many client-side libraries and applications are silently ignoring
// all non-important headers, which would ultimately hide such warning from the
// actual client code logic, unless the programmer directly checks the known
// header.
//
// But in cases when the server responds with "267 Questionable But Okay", the
// request status code becomes obvious. There are a lot of implementations that
// expect the response code to be strictly equal to "200 OK" and thus would
// consider the potentially erroneous request as failed indeed, bringing
// developers' attention to the underlying problem.
//
// From the practical point of view, the server may return 267 status code for
// GET requests each time it determines that the particular request is not
// coming from an official or fully conforming REST client. Small differences in
// implementation of client logic might make this client application to stand
// out from the behavior patterns of official clients:
//
//   - Absence of a GET parameter that is always sent by official clients – in
//     cases when the server can tolerate its absence and still fulfil the
//     request based on other available information
//     (e.g. "/getPhoto?album_id=421&photo_index=NaN" might fallback to photo_index=0)
//   - Duplication of a GET parameter or receiving GET parameters in an
//     unexpected order – in cases where the order of those parameters has a
//     semantic meaning for this particular API endpoint and the server logic
//     has valid reasons to specifically check for this
//     (e.g. "/getPhoto?photo_index=3&album_id=421&photo_index=3")
//   - Presence of a GET parameter that has a special meaning for other API
//     endpoints – in cases when this parameter otherwise involves a different
//     processing on the server that is not required for this particular
//     endpoint (e.g. "/getPhoto?album_id=421&photo_index=3&user_token=abcdef")
//   - Presence of a HTTP header that has a special meaning for other API
//     endpoints and requires a special processing that cannot be applied for
//     this particular request (e.g. "Authorization: Basic YmFzZTY0")
//
// This status code should be related only to this particular requests,
// indicating a discrepancy in logic that formed it. It should not be used to
// indicate a questionable series of requests made by client (e.g. wrong order
// of API calls).
// Also it is not advisable to use it with POST requests that return an
// application-defined data structure, since it is naturally to include any
// warnings or hints directly in the POST body.
//
// Generally, a received "267 Questionable But Okay" status code indicates that
// the server warns the client that the request is identified as non-genuine but
// otherwise conforming to web standards. It may be silently ignored or
// converted to "200 OK" on the client side in case the programmer
// deliberatively wants to simplify the code and generalize the interactions
// with server API.
//
// However, its highly advisable to log all 267 responses in fully conforming
// client because it indicates a potential bug in request logic. A conforming
// client should be prepared to receive 267 statuses for any REST call to the
// server that is known to issue them, but must not treat them as fatal errors.
//
// A third-party client that had strict checks for equality to status code 200
// would essentially fail the questionable request, bringing developers'
// attention to the fact that the logic could be buggy, and the code should be
// modified according to changes in official conforming applications, or updated
// to ignore them, accepting the fact that the server is not happy.
//
// If a request resulted in 267 response, then any subsequent request with the
// same parameters should also return 267 status code unless the server state
// was changed in a way that even a genuine request would have been resulted in
// a different status code too.
//
// See: https://github.com/maximal/http-267/issues/1#issuecomment-2000395105
//
// 267 code can be used as a soft deprecation warning for this particular API
// endpoint and its passed or omitted parameters. Example: API version ?v= is
// not passed and defaults to the latest version before addition of explicit
// versioning, that is expected to be shut down soon.
//
// If the updated server decides that a particular cause of 267 is no longer can
// be tolerated, it may start to return "422 Unprocessable Entity" instead, as a
// refusal to support this deprecated API call.
//
// For example: Notification for the client that performs some extra work which
// is not required for this particular API call (e.g. it sends a header with
// cryptographically signed URL parameters even for requests that do not need to
// be signed, contrary to other API calls that may require it).
//
// See: https://github.com/maximal/http-267/issues/1#issuecomment-2000599066
const StatusQuestionableButOkay = 267
