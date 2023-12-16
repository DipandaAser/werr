// create a map with a list of routes
export const routes = {
    HOME: "/",
    LOGIN: "/login",
    ACCOUNT: {
        BASE: "/account",
        MEDIA: {
            UPLOAD: "/account/media/upload",
        }
    }
}


export function ToTitleCase(str: string) {
    return str.charAt(0).toUpperCase() + str.slice(1).toLowerCase();
}