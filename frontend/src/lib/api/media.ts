import type { Axios } from "axios"
import { NewApiClient } from "./apiClient"
import type { AxiosInstance } from "axios"

const baseURL = '/media'
export class MediaAPI {
    private axiosInstance: AxiosInstance

    constructor(axiosInstance: AxiosInstance) {
        this.axiosInstance = axiosInstance
    }

    ReqUpload(name: string, file: File) {
        return new Promise<MediaUploadResponse>((resolve, reject) => {
            return this.axiosInstance.post(
                `${baseURL}?title=${encodeURIComponent(name)}`,
                file)
                .then(repsonse => {
                    const data: MediaUploadResponse = repsonse.data
                    resolve(data)
                })
                .catch(err => {
                    reject(err.response.data.error)
                })
        })
    }
}

export interface MediaUploadResponse extends Media {

}

export enum MediaStatus {
    Processing = "processing",
    Completed = "completed",
    Draft = "draft"
}

export enum MediaType {
    Image = "image",
    Video = "video",
    Audio = "audio",
    Document = "document",
    Other = "other"
}


export interface Media {
    id: string;
    created_at: Date;
    title: string;
    status: MediaStatus;
    type: MediaType;
    tags: string[];
    user_id: string;
}