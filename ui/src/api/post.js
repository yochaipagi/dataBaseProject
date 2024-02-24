
export default function post(url, data, resetForm, setIsLoading) {
    return fetch(url, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(data)
    })
        .then(res => res.json())
        .then(
            (data) => {
                if (data.error) {
                    alert(`An error occurred: ${data.error}`)
                } else {
                    alert("completed successfully")
                    resetForm()
                }
            },
            (err) => alert(`An error occurred: ${err.toString()}`)
        )
        .finally(() => setIsLoading(false))
};
