const PUBLIC_URL = process.env.PUBLIC_URL || process.env.REACT_APP_API_HOST || '';

export const getImagePath = (imageUrl: string) => {
    return `${PUBLIC_URL}/${imageUrl}`
}