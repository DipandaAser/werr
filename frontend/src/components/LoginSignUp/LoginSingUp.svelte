<script lang="ts">
  import Icon from "@iconify/svelte";
  import { authPopupStore, closeAuthPopup } from "../../stores/authPopupStore";
  import { authHandlers } from "../../stores/authStore";
  import { getFirebaseAuthErrorMessage } from "$lib/firebase/firebase.client";
  import LoadingButton from "./LoadingButton.svelte";
  export let register = true;
  let passwordVisibility = false;
  let email = "";
  let password = "";
  let username = "";
  let isLoading = false;

  //error
  let showError = false;
  let errorMessage = "";

  const togglePasswordVisibility = () => {
    passwordVisibility = !passwordVisibility;
    const pwdInput = document.getElementById("password") as HTMLInputElement;
    if (pwdInput.type === "password") {
      pwdInput.type = "text";
    } else {
      pwdInput.type = "password";
    }
  };

  async function join(isGoogle = false) {
    clearError();
    try {
      if (isGoogle) {
        await authHandlers.joinWithGoogle();
      } else if (register) {
        isLoading = true;
        await authHandlers.signup(username, email, password);
      } else {
        isLoading = true;
        await authHandlers.login(email, password);
      }

      isLoading = false;
      closeAuthPopup();
      return;
    } catch (error) {
      isLoading = false;
      const errMessage = getFirebaseAuthErrorMessage(error as { code: string });
      console.log("error message: ", errorMessage);
      showError = true;
      errorMessage = errMessage;
    }
  }

  function clearError() {
    showError = false;
    errorMessage = "";
  }

  register = $authPopupStore.isRegisterMode;
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div>
  <div class="backdrop">
    <div class="container">
      <div class="content">
        <div class="promo">
          Sign up to download unlimited full resolution media
        </div>
        <div class="authPanel">
          <div class="scrollingPanel">
            <div class="scrollingPanel-container">
              <div class="scrollingPanel-container2">
                <div class="handles">
                  <button
                    on:click={() => (register = true)}
                    class="handle {register ? 'active' : ''}">Register</button
                  >
                  <button
                    on:click={() => (register = false)}
                    class="handle {!register ? 'active' : ''}">Login</button
                  >
                </div>
                <div class="panel">
                  <form
                    id={register ? "signupForm" : "loginForm"}
                    action={register ? "signup" : "login"}
                    class="panel-form-container"
                  >
                    <div class="socialAuthButtons">
                      <div class="socialAuthButtons-container">
                        <button on:click={() => join(true)} type="button"
                          ><svg
                            viewBox="0 0 24 24"
                            xmlns="http://www.w3.org/2000/svg"
                            fill-rule="evenodd"
                            clip-rule="evenodd"
                            stroke-linejoin="round"
                            stroke-miterlimit="2"
                            ><path
                              d="M0 12C0 5.373 5.373 0 12 0s12 5.373 12 12-5.373 12-12 12S0 18.627 0 12z"
                              fill="#f7f7f7"
                            ></path><path
                              d="M8.659 13.252l-.417 1.559-1.527.033A5.98 5.98 0 016 12c0-.995.242-1.933.671-2.759l1.359.249.596 1.351A3.554 3.554 0 008.433 12c0 .441.08.863.226 1.252z"
                              fill="#d14848"
                              fill-rule="nonzero"
                            ></path><path
                              d="M17.895 10.879a6.001 6.001 0 01-.026 2.374 5.996 5.996 0 01-2.113 3.426l-1.712-.088-.242-1.512a3.584 3.584 0 001.538-1.826h-3.208v-2.374h5.763z"
                              fill="#518ef8"
                              fill-rule="nonzero"
                            ></path><path
                              d="M15.756 16.679a5.998 5.998 0 01-9.041-1.836l1.944-1.591a3.568 3.568 0 005.142 1.827l1.955 1.6z"
                              fill="#00ab6b"
                              fill-rule="nonzero"
                            ></path><path
                              d="M15.83 7.381l-1.944 1.592a3.568 3.568 0 00-5.26 1.868l-1.955-1.6A5.997 5.997 0 0112 6c1.456 0 2.791.519 3.83 1.381z"
                              fill="#d14848"
                              fill-rule="nonzero"
                            ></path></svg
                          ><span>Continue with Google</span></button
                        >
                      </div>
                    </div>
                    <div class="separator">
                      <hr />
                      <div><span>or</span></div>
                    </div>
                    {#if register}
                      <div class="inputField">
                        <!-- svelte-ignore a11y-label-has-associated-control -->
                        <label for="username">* Username</label>
                        <div class="textInputContainer">
                          <div class="textInputField">
                            <input
                              id="username"
                              name="username"
                              class="textInput"
                              type="text"
                              placeholder="example maria93"
                              bind:value={username}
                            />
                          </div>
                        </div>
                      </div>
                      <div class="inputField">
                        <!-- svelte-ignore a11y-label-has-associated-control -->
                        <label for="email">* Email address</label>
                        <div class="textInputContainer">
                          <div class="textInputField">
                            <input
                              id="email"
                              name="email"
                              class="textInput"
                              type="email"
                              placeholder="example@gmail.com"
                              bind:value={email}
                            />
                          </div>
                        </div>
                      </div>
                      <div class="passwordContainer">
                        <!-- svelte-ignore a11y-label-has-associated-control -->
                        <label for="password">* Password</label>
                        <div class="textInputContainer">
                          <div class="textInputField">
                            <input
                              id="password"
                              name="password"
                              type="password"
                              class="textInput padRight"
                              bind:value={password}
                            />
                            <div class="addonAfter">
                              <button
                                type="button"
                                class="togglePasswordButton"
                                on:click={togglePasswordVisibility}
                                aria-label="Toggle password visible"
                                ><Icon
                                  icon={passwordVisibility
                                    ? "basil:eye-closed-solid"
                                    : "basil:eye-solid"}
                                  class="icon"
                                  width="24"
                                  height="24"
                                ></Icon></button
                              >
                            </div>
                          </div>
                        </div>
                      </div>
                      <div>
                        <div>
                          <div>
                            <div
                              role="progressbar"
                              aria-valuenow="0"
                              aria-valuemin="0"
                              aria-valuemax="4"
                            ></div>
                          </div>
                          <div>
                            Strengh<span>&nbsp weak</span>
                          </div>
                        </div>
                      </div>
                      <div>*&nbsp;at least 8&nbsp;char, inclusing 1 number</div>
                    {:else}
                      <div class="inputField">
                        <!-- svelte-ignore a11y-label-has-associated-control -->
                        <label for="email">* Email address</label>
                        <div class="textInputContainer">
                          <div class="textInputField">
                            <input
                              id="email"
                              name="email"
                              class="textInput"
                              type="email"
                              placeholder="example@gmail.com"
                              bind:value={email}
                            />
                          </div>
                        </div>
                      </div>
                      <div class="passwordContainer">
                        <!-- svelte-ignore a11y-label-has-associated-control -->
                        <label for="password">* Password</label>
                        <div class="textInputContainer">
                          <div class="textInputField">
                            <input
                              id="password"
                              name="password"
                              type="password"
                              class="textInput padRight"
                              bind:value={password}
                            />
                            <div class="addonAfter">
                              <button
                                type="button"
                                class="togglePasswordButton"
                                on:click={togglePasswordVisibility}
                                aria-label="Toggle password visible"
                                ><Icon
                                  icon={passwordVisibility
                                    ? "basil:eye-closed-solid"
                                    : "basil:eye-solid"}
                                  class="icon"
                                  width="24"
                                  height="24"
                                ></Icon></button
                              >
                            </div>
                          </div>
                        </div>
                      </div>
                    {/if}
                    {#if showError}
                      <div class="error">
                        {errorMessage}
                      </div>
                    {/if}

                    <div class="recaptchaExplainer">
                      This site is protected by reCAPTCHA and the Google <a
                        href="https://policies.google.com/privacy"
                        >Privacy Policy</a
                      >
                      and
                      <a href="https://policies.google.com/terms"
                        >Terms of Service</a
                      > apply.
                    </div>
                    <div class="submitContainer">
                      <button
                        class="joinButton"
                        type="submit"
                        on:click={(e) => {
                          e.preventDefault();
                          join(false);
                        }}
                      >
                        {#if isLoading}
                          <LoadingButton />
                        {:else}
                          <span class="label"
                            >{register ? "Create account" : "Login"}</span
                          >
                        {/if}
                      </button>
                    </div>
                  </form>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <button
        on:click={closeAuthPopup}
        class="closeButton"
        aria-label="Close"
        aria-keyshortcuts="Escape"
      >
        <Icon icon="basil:cross-solid" width="36" height="36" />
      </button>
    </div>
  </div>
</div>

<style lang="scss">
  $green: #00ab6b;
  $hoverColor: rgba(182, 175, 175, 0.151);
  .backdrop {
    height: 100%;
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    height: calc(var(--vh, 1vh) * 100);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 10;
    background-color: rgba(25, 27, 38, 0.3);
  }

  .error {
    color: #d14848;
    font-size: 14px;
    font-weight: bolder;
    margin-top: 16px;
  }

  .container {
    position: relative;
    height: 692px;
    max-height: 90vh;
  }

  .content {
    border-radius: 0;
    height: 100%;
    display: flex;
    align-items: flex-start;
    border-radius: 4px;
    overflow: hidden;
    flex-direction: column;
    justify-content: flex-start;
  }

  .promo {
    max-width: 434px;
    width: 100vw;
    flex: initial;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #191b26;
    text-align: center;
    color: #fff;
    padding: 16px;
    font-weight: 800;
    font-size: 16px;
    padding-top: 24px;
  }

  .authPanel {
    flex: auto;
    overflow: hidden;
    flex-shrink: 1;
    max-width: 434px;
    width: 100vw;
    background-color: #fff;
    position: relative;
  }

  .scrollingPanel {
    overflow-y: auto;
    min-height: inherit;
    height: 100%;
    width: 100%;

    .scrollingPanel-container {
      height: auto;
      padding-top: 16px;
      padding-bottom: 0;
      background: #fff;
      width: 100%;
      padding-left: 24px;
      padding-right: 24px;

      .scrollingPanel-container2 {
        width: 100%;

        .handles {
          padding: 0 16px;
          width: 100%;
          display: flex;
          align-items: flex-end;
          margin-bottom: 16px;

          .handle {
            border: none;
            background: none;
            padding: 0;
            flex: 1;
            font-size: 16px;
            font-weight: 600;
            color: #191b26;
            cursor: pointer;
            margin-right: 16px;
            margin-bottom: 8px;
            padding-bottom: 8px;
            border-bottom: 2px solid transparent;
            transition: border-bottom 0.2s;

            &.active {
              border-bottom: 3px solid $green;
              color: $green;
            }

            &:hover {
              border-bottom: 3px solid $green;
              color: $green;
              background-color: $hoverColor;
            }
          }
        }

        .panel {
          width: 100%;
          .panel-form-container {
            width: 100%;
            display: flex;
            flex-direction: column;
            padding: 16px 0 0;
            min-height: inherit;
            .socialAuthButtons {
              display: flex;
              width: 100%;
              flex-direction: column;

              .socialAuthButtons-container {
                display: flex;
                flex-direction: column;

                button {
                  position: relative;
                  width: 100%;
                  border-radius: 24px;
                  height: 40px;
                  display: flex;
                  align-items: center;
                  justify-content: center;
                  border: 1px solid #eaeaea;
                  padding: 0;
                  background: #fff;
                  font-weight: 600;
                  color: #191b26;
                  cursor: pointer;
                  appearance: button;
                  font-size: 14px;
                  overflow: visible;
                  text-transform: none;
                  margin: 0;

                  &:hover {
                    background-color: #0060df;
                    color: #fff;
                  }

                  svg {
                    width: 24px;
                    height: 24px;
                    position: absolute;
                    left: 8px;
                    top: 7px;
                    overflow: hidden;
                    fill: currentColor;
                  }
                }
              }
            }

            .separator {
              position: relative;
              height: 24px;
              width: 100%;
              margin: 24px 0 8px;

              hr {
                position: absolute;
                top: 12px;
                width: 100%;
                padding: 0;
                margin: 0;
                border-color: #eaeaea;
                border-top: 0 #eaeaea;
                border-style: solid;
                color: inherit;
                height: 0;
                overflow: visible;
              }

              div {
                position: absolute;
                left: 0;
                top: 0;
                height: 24px;
                width: 100%;
                display: flex;
                align-items: center;
                justify-content: center;

                span {
                  font-size: 14px;
                  font-weight: 600;
                  color: #656f79;
                  background: #fff;
                  padding: 0 16px;
                  line-height: 24px;
                }
              }
            }
          }
        }
      }
    }
  }

  .inputField {
    margin-bottom: 16px;
    margin-top: 16px;

    label {
      display: flex;
      font-weight: 600;
      margin-bottom: 8px;
      font-size: 14px;
    }
  }

  .textInputContainer {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    width: 100%;
  }

  .textInputField {
    position: relative;
    width: 100%;

    .textInput {
      width: 100%;
      border: 1px solid #eaeaea;
      border-radius: 4px;
      min-height: 40px;
      display: flex;
      align-items: center;
      padding: 0 12px;
      color: #191b26;
      font-size: 14px;
      font-weight: 400;
      overflow: visible;
      margin: 0;

      &:focus {
        outline: 1px #00ab6b;
        border-color: #00ab6b;
      }

      &.padright {
        padding-right: 40px;
      }
    }
  }

  .addonAfter {
    right: 8px;
    width: 24px;
    height: 24px;
    position: absolute;
    top: 8px;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #656f79;

    .togglePasswordButton {
      -webkit-appearance: button;
      appearance: button;
      background: none;
      border: none;
      display: flex;
      align-items: center;
      justify-content: center;
      color: #656f79;
    }
  }

  .passwordContainer {
    margin-top: 16px;

    label {
      display: flex;
      font-weight: 600;
      margin-bottom: 8px;
      font-size: 14px;
    }
  }

  .closeButton {
    border-style: none;
    padding: 0;
    right: 16px;
    top: 8px;
    position: absolute;
    display: flex;
    right: -0px;
    top: 0;
    align-items: center;
    justify-content: center;
    background: none;
    border: none;
    border-radius: 4px;
    padding: 0;
    font-size: 20px;
    color: #fff;
    cursor: pointer;
  }

  .recaptchaExplainer {
    color: #424141;
    font-size: 12px;
    margin-top: 16px;

    a {
      color: #00ab6b;
      text-decoration: none;
    }
  }

  .submitContainer {
    width: 100%;
    position: sticky;
    bottom: 0;
    padding: 8px 0 24px;
    background-color: #fff;
    margin-top: 16px;

    .joinButton {
      width: 100%;
      border: none;
      background: #00ab6b;
      color: #fff;
      -webkit-appearance: button;
      appearance: button;
      cursor: pointer;
      border-radius: 24px;
      height: 40px;
      display: inline-flex;
      align-items: center;
      justify-content: center;
      padding: 0 32px;
      min-width: 160px;
      position: relative;
      transition: all 0.1s ease-in;
      font-size: 14px;
      font-weight: 400;
      overflow: visible;
      text-transform: none;
      margin: 0;

      .label {
        font-weight: 600;
        color: inherit;
        cursor: pointer;
      }
    }
  }
</style>
