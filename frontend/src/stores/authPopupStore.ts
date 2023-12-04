import { writable } from "svelte/store"

export const authPopupStore = writable<{
    isOpen: boolean;
    isRegisterMode: boolean;
}>({
    isOpen: false,
    isRegisterMode: false,
});

export const openAuthPopup = (isRegisterMode = false) => {
    authPopupStore.update((store) => ({
        ...store,
        isOpen: true,
        isRegisterMode: isRegisterMode,
    }));
}

export const closeAuthPopup = () => {
    authPopupStore.update((store) => ({
        ...store,
        isOpen: false,
    }));
}

export const toggleAuthPopup = () => {
    authPopupStore.update((store) => ({
        ...store,
        isOpen: !store.isOpen,
    }));
}