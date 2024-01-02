import { env } from '$env/dynamic/public'
import axios from 'axios'
import { authStore } from '../../stores/authStore'
import type { Axios } from 'axios'
import type { AxiosInstance } from 'axios'
import { MediaAPI } from './media'

export function NewApiClient(token: string = ""): ApiClient {
    const axiosClient = axios.create({
        baseURL: env.PUBLIC_API_URL,
        headers: {
            Token: token
        }
    })

    const apiClient = new ApiClient(axiosClient)

    return apiClient
}

export class ApiClient {
    axiosInstance: AxiosInstance
    Media: MediaAPI

    constructor(axiosInstance: AxiosInstance) {
        this.axiosInstance = axiosInstance
        this.Media = new MediaAPI(axiosInstance)
    }
}