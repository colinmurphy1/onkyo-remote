// getStatus gets the status of the receiver from the /api/status endpoint.
const getStatus = async () => {
    // Make an http request
    const req = await fetch("/api/status", {
        method: "GET",
    })
    .then(response => {
        if (! response.ok) {
            throw new Error("BadResponseError")
        }
        return response.json()
    })
    .then(data => data)

    return req.data
}

export default getStatus