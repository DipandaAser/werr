import { writable } from "svelte/store";
import { auth } from "$lib/firebase/firebase.client";
import { GoogleAuthProvider, createUserWithEmailAndPassword, getRedirectResult, sendPasswordResetEmail, signInWithEmailAndPassword, signInWithPopup, signInWithRedirect, signOut, updateEmail, updatePassword } from "firebase/auth";
import type { User as FirebaseUser } from "firebase/auth";
import { closeAuthPopup } from "./authPopupStore";

export const authStore = writable<{
    isLoading: boolean;
    currentUser: FirebaseUser | null;
}>({
    isLoading: true,
    currentUser: null,
});

export const authHandlers = {
    login: async (email: string, password: string) => {
        const userCredential = await signInWithEmailAndPassword(auth, email, password)
        authStore.update(store => {
            store.currentUser = userCredential.user;
            store.isLoading = false;
            return store;
        })
    },
    signup: async (username: string, email: string, password: string) => {
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
    joinWithGoogle: async () => {
        console.log("joinWithGoogle");
        const provider = new GoogleAuthProvider();
        provider.addScope('profile');
        provider.addScope('email');
        await signInWithPopup(auth, provider)
        const result = await getRedirectResult(auth);
        if (result) {
            //const credential = GoogleAuthProvider.credentialFromResult(result);
            //const token = credential?.accessToken;
            authStore.update(store => {
                console.log("joinWithGoogle", result);
                store.currentUser = result.user;
                store.isLoading = false;
                return store;
            })
        }
    }
}