/**
 * Enum for available media categories
 *
 * @export @readonly
 * @enum {string}
 */
export enum Categories {
    IMAGES = 'images',
    GIFS = 'gifs',
    ILLUSTRATIONS = 'illustrations',
    VIDEOS = 'videos',
    AUDIOS = 'audios',
}

export enum CategoriesLabels {
    IMAGES = 'Images',
    GIFS = 'Gifs',
    ILLUSTRATIONS = 'Illustrations',
    VIDEOS = 'Videos',
    AUDIOS = 'Audios',
}

export enum CategoriesIcons {
    IMAGES = 'basil:camera-solid',
    GIFS = 'basil:fire-solid',
    ILLUSTRATIONS = 'material-symbols:draw-outline-rounded',
    VIDEOS = 'basil:video-solid',
    AUDIOS = 'basil:music-solid',
}

export enum CategoriesUrls {
    IMAGES = '/images',
    GIFS = '/gifs',
    ILLUSTRATIONS = '/illustrations',
    VIDEOS = '/videos',
    AUDIOS = '/audios',
}

export const CategoriesList = [
    {
        label: CategoriesLabels.IMAGES,
        icon: CategoriesIcons.IMAGES,
        value: Categories.IMAGES,
        url: CategoriesUrls.IMAGES,
    },
    {
        label: CategoriesLabels.GIFS,
        icon: CategoriesIcons.GIFS,
        value: Categories.GIFS,
        url: CategoriesUrls.GIFS,
    },
    {
        label: CategoriesLabels.VIDEOS,
        icon: CategoriesIcons.VIDEOS,
        value: Categories.VIDEOS,
        url: CategoriesUrls.VIDEOS,
    },
    {
        label: CategoriesLabels.ILLUSTRATIONS,
        icon: CategoriesIcons.ILLUSTRATIONS,
        value: Categories.ILLUSTRATIONS,
        url: CategoriesUrls.ILLUSTRATIONS,
    },

    {
        label: CategoriesLabels.AUDIOS,
        icon: CategoriesIcons.AUDIOS,
        value: Categories.AUDIOS,
        url: CategoriesUrls.AUDIOS,
    }
];

export function isCategoriesURLHome(url: string): boolean {
    console.log("THE URL: ", url);
    if (url == '/') {
        return true;
    }
    // check if url is one of the categories url
    // check if the category url is on the form of /[language]/[category]
    const regexHome = new RegExp(`^\/([a-z]{2})$`, 'g');
    return CategoriesList.some((category) => {
        if (category.url == url) {
            return true
        }
        //const regex = new RegExp(`^(\/([a-z]{2})?)?\/${category.value}`, 'i');
        const regex = new RegExp(`^((\/([a-z]{2})?)?\/${category.value}$)|(\/([a-z]{2})?$)`, 'gi');
        return regex.test(url);
    });
}