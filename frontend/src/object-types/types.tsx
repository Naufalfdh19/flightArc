export interface Err {
    field: string
    message: string
}

export interface Base<T> {
    data: T
    message: string
    error: Err[]
}