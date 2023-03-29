import { useSetRecoilState } from "recoil";
import { useFetchWrapper } from "../helpers";

export const useImageUploadActions = () => {
    const baseApiUrl = `${process.env.REACT_APP_API_URL}/upload`;
    const baseUrl = `${process.env.REACT_APP_API_HOST}`;
    const fetchWrapper = useFetchWrapper();

    const upload = async (image: File) => {
        let formData = new FormData();
        formData.append('image', image, image.name);
        const r = await fetchWrapper.post(baseApiUrl, formData);
        return r.data;
    }
    
    const download = async (url: string) => {
        const r = await fetch(`${baseUrl}/${url}`);
        return await r.blob();
    }

    return {
        upload,
        download,
    }
}