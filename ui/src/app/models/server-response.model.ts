export interface ServerResponse<T> {
  ok: boolean
  data?: T
  error?: string

}
