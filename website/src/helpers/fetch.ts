export const useFetchWrapper = () => {
    const request = (method: string) => {
        return async (url: string, body: object | null) => {
            const requestOptions: RequestInit = {
                method,
                headers: {
                    'Content-Type': 'application/json'
                },
                body,
            } as RequestInit;

            if (body) {
                requestOptions.body = JSON.stringify(body);
            }
            const response = await fetch(url, requestOptions);
            return handleResponse(response);
        }
    }

    
    const handleResponse = async (response: Response) => {
        const text = await response.text();
        const data = text && JSON.parse(text);
        if (!response.ok) {
            const error = (data && data.message) || response.statusText;
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