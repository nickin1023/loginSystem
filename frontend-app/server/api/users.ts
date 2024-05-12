import type { IncomingMessage, ServerResponse } from 'http'
import axios from 'axios'

export default async (req: IncomingMessage, res: ServerResponse) => {
  
  const axiosInstance = axios.create({
    baseURL: 'http://go:8000'
  })

  try {
    const response = await axiosInstance.get('/admin/users')

    return {
      response: response.data
    }
  } catch (e) {
    console.log('be communicate error.')
  }
}