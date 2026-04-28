import { useCallback, useEffect, useState } from "react";
import type { Err } from "../object-types/types";


export  default function useFetch<tData>(
    url: string
) {
    const [data, setData] = useState<tData>()
    const [error, setError] = useState<Err>()
    const [isLoading, setIsLoading] = useState(true)


    const fetchData = useCallback(
        async (options?: RequestInit) => {
            setError(undefined)
            
            const res = await fetch(
                url, options
            );
            
            const resData = await res.json()

            setData(resData)
            
            setIsLoading(false);
            
        }, [url])

        useEffect(() => {
            fetchData();
        }, [fetchData])

    return {data, error, isLoading, fetchData}
}