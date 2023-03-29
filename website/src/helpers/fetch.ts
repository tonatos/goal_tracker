export const useFetchWrapper = () => {
    const request = (method: string) => {
        return async (url: string, body?: FormData | object, headers?: object) => {
            const requestOptions: RequestInit = {
                method,
                headers: headers ? headers : {},
                body,
            } as RequestInit;

            if (!(body instanceof FormData)) {
                requestOptions.headers = {
                    ...requestOptions.headers,
                    'Content-Type': 'application/json',
                }
            }

            if (body) {
                requestOptions.body = body instanceof FormData ? body : JSON.stringify(body);
            }
            
            const response = await fetch(url, requestOptions);
            return handleResponse(response);
        }
    }

    
    const handleResponse = async (response: Response) => {
        const text = await response.text();
        const data = text && JSON.parse(text);
        if (!response.ok) {
            // if (response.status < 500) {
            //     return data;
            // }

            const error = data || response.statusText;
            return Promise.reject(error);
        }
        return data;
    }

    return {
        get: request('GET'),
        post: request('POST'),
        put: request('PUT'),
        delete: request('DELETE')
    };
}