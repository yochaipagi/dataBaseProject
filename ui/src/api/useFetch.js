import {useEffect, useState} from "react";

export default function useFetch(url, options, expect={}, deps) {
    const [response, setResponse] = useState(expect);
    const [error, setError] = useState(null);
    const [isLoading, setIsLoading] = useState(false);

    useEffect(() => {
        fetch(url, {...options})
            .then(res => res.json())
            .then(
                (res) => {
                    setResponse(res.response);
                    setIsLoading(false);
                }, (error) => {
                    setError(error)
                    setIsLoading(false);
                }
            )
    }, [url, options, deps]);

    return [response, isLoading, error];
};
