import { axios } from "@/lib/axios"

export type AllFields = {
  fields: string[]
}

export const getAllFields = (): Promise<AllFields> => {
  return axios.get('/fields');
}
