import { writable } from "svelte/store";
import { auth } from "$lib/firebase/firebase.client";
import { createUserWithEmailAndPassword, sendPasswordResetEmail, signInWithEmailAndPassword, signOut, updateEmail, updatePassword } from "firebase/auth";
import type { User as FirebaseUser } from "firebase/auth";

export const authStore = writable<{
    isLoading: boolean;
    currentUser: FirebaseUser | null;
}>({
    isLoading: true,
    currentUser: null,
});

export const authHandlers = {
    login: async (email: string, password: string) => {
        await signInWithEmailAndPassword(auth, email, password)
    },
    signup: async (email: string, password: string) => {
        const userCredential = await createUserWithEmailAndPassword(auth, email, password);
        authStore.update(store => {
            store.currentUser = userCredential.user;
            store.isLoading = false;
            return store;
        })
    },
    logout: async () => {
        await signOut(auth);
    },
    resetPassword: async (email: string) => {
        await sendPasswordResetEmail(auth, email);
    },
    updateEmail: async (email: string) => {
        if (!auth.currentUser) throw new Error("No user logged in");

        await updateEmail(auth.currentUser, email);
    },
    updatePassword: async (password: string) => {
        if (!auth.currentUser) throw new Error("No user logged in");

        await updatePassword(auth.currentUser, password);
    },
}