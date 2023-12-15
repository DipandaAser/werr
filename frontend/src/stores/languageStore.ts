import { sourceLanguageTag, type AvailableLanguageTag } from "$paraglide/runtime";
import { writable } from "svelte/store"

export const languageStore = writable<{
    lang: AvailableLanguageTag;
}>({
    lang: sourceLanguageTag,
});

export const updateLang = (lang: AvailableLanguageTag) => {
    languageStore.update((store) => ({
        ...store,
        lang: lang,
    }));
}