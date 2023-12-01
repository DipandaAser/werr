// create a map with a list of routes

export const routes = {
    HOME: "/",
    ACCOUNT: "/account",
    LOGIN: "/login",
}


export function ToTitleCase(str: string) {
    return str.charAt(0).toUpperCase() + str.slice(1).toLowerCase();
}