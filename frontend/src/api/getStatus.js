// getStatus gets the status of the receiver from the /api/status endpoint.
const getStatus = async () => {
    // Make an http request
    const req = await fetch("/api/status", {
        method: "GET",
    })
    .then(response => response.json())
    .then(data => data)

    // Verify response is 200 OK
    /*if (req.status != 200) {
        throw "BadResponseError"
    }*/

    return req.data
}

export default getStatus