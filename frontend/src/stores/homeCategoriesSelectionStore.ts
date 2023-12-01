import { writable } from "svelte/store";
import { searchStore, updateSearchStore } from "./searchStore";

export const chipsStore = writable<{
    selected: string;
}>({
    selected: "",
});


export const updateSelectedChip = (selected: string) => {
    chipsStore.update(store => {
        store.selected = selected;
        return store;
    })
}