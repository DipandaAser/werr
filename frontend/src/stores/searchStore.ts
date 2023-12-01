import { writable } from "svelte/store";

export const searchStore = writable<{
    value: string;
    categorie: string;
}>({
    value: "",
    categorie: "",
});

export const updateSearchStore = (value: string, categorie: string) => {
    searchStore.update(store => {
        store.value = value;
        store.categorie = categorie;
        return store;
    })
}