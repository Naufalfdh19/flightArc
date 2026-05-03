import { useCallback, useEffect, useState } from "react";
import type { Err } from "../object-types/types";


export  default function useFetch<tData>(
    url: string
) {
    const [data, setData] = useState<tData>()
    const [error, setError] = useState<Err | Error>()
    const [isLoading, setIsLoading] = useState(false)


    const fetchData = useCallback(
        async (options?: RequestInit) => {
            setIsLoading(true)
            setError(undefined)
            
            try {
                const res = await fetch(
                    url, options
                 );
            
                const resData = await res.json()
                setData(resData)

            }  catch (err) {
                setError(err instanceof Error ? err : new Error(String(err)));
            } finally {
                setIsLoading(false);
            }

            
        }, [url])

    useEffect(() => {
        fetchData();
    }, [fetchData])

    return {data, error, isLoading, fetchData}
}