import { writable } from "svelte/store"

export const accountDropdownStore = writable<{
    isOpen: boolean;
}>({
    isOpen: false,
});

export const openDropDown = () => {
    accountDropdownStore.update((store) => ({
        ...store,
        isOpen: true,
    }));
}

export const closeDropDown = () => {
    accountDropdownStore.update((store) => ({
        ...store,
        isOpen: true,
    }));
}