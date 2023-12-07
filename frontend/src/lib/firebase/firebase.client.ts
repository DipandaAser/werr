// Import the functions you need from the SDKs you need
import { deleteApp, getApp, getApps, initializeApp } from "firebase/app";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries
import { getAuth, setPersistence, inMemoryPersistence, type AuthError } from 'firebase/auth'
import { env } from '$env/dynamic/public';

// Your web app's Firebase configuration
const firebaseConfig = {
    apiKey: env.PUBLIC_FIREBASE_API_KEY,
    authDomain: env.PUBLIC_FIREBASE_AUTH_DOMAIN,
    appId: env.PUBLIC_FIREBASE_APP_ID
};

// Initialize Firebase
let firebaseApp;
if (!getApps().length) {
    firebaseApp = initializeApp(firebaseConfig)
} else {
    firebaseApp = getApp()
    deleteApp(firebaseApp)
    firebaseApp = initializeApp(firebaseConfig)
}

export function getFirebaseAuthErrorMessage(error: { code: string }): string {
    let message = 'An error occurred. Please try again.';

    switch (error.code) {
        case 'auth/invalid-email':
            message = 'The email address is invalid.';
            break;
        case 'auth/user-disabled':
            message = 'This user has been disabled. Please contact support.';
            break;
        case 'auth/user-not-found':
            message = 'No user found with this email address.';
            break;
        case 'auth/wrong-password':
            message = 'Incorrect password. Please try again.';
            break;
        case 'auth/email-already-in-use':
            message = 'This email is already in use. Please use a different email.';
            break;
        case 'auth/weak-password':
            message = 'The password is too weak. Please use a stronger password.';
            break;
        case 'auth/too-many-requests':
            message = 'Too many requests. Please try again later.';
            break;
        case 'auth/invalid-login-credentials':
            message = 'Invalid login credentials. Please try again.';
            break;
        // Add more cases as needed for different Firebase Auth errors
        default:
            message = 'An error occurred. Please try again.';
            console.error('Firebase Auth Error:', error);
            break;
    }

    return message;
}

export const auth = getAuth(firebaseApp)

