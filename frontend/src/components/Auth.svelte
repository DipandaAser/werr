<script lang="ts">
  import { routes } from "$lib";
  import { authHandlers, authStore } from "../stores/authStore";

  let register = false;
  let displayError = false;
  let errorMessage = "";
  let email = "";
  let password = "";
  let confirmPassword = "";

  async function handleSubmit() {
    if (!email || !password || (register && !confirmPassword)) {
      return;
    }

    if (register && password === confirmPassword) {
      try {
        await authHandlers.signup(email, password);
      } catch (err) {
        console.log("signup: ", err);
      }
    } else {
      try {
        await authHandlers.login(email, password);
      } catch (err) {
        // display an error message in the UI
        console.log("login err: ", err);
        // displayError = true;
        // errorMessage = (err as { message: string }).message;
      }
    }
    if ($authStore.currentUser) {
      window.location.href = routes.ACCOUNT;
    }
  }

  $: console.log("displayError", displayError);
</script>

<div class="container">
  <h1>{register ? "Register" : "Log in"}</h1>
  <form>
    <label>
      <input
        on:change={() => {
          displayError = false;
        }}
        bind:value={email}
        type="email"
        placeholder="Email"
      />
    </label>
    <label>
      <input bind:value={password} type="password" placeholder="Password" />
    </label>
    {#if register}
      <label>
        <input
          bind:value={confirmPassword}
          type="password"
          placeholder="Confirm Password"
        />
      </label>
    {/if}
    <button on:click={handleSubmit}>Submit</button>
  </form>
  {#if register}
    <div
      role="button"
      tabindex="0"
      on:click={() => {
        register = false;
      }}
      on:keydown={() => {}}
    >
      Already have an account?
      <p>Log in</p>
    </div>
  {:else}
    <div
      role="button"
      tabindex="0"
      on:click={() => {
        register = true;
      }}
      on:keydown={() => {}}
    >
      Don't have an account? <p>Sign Up</p>
    </div>
    <div
      role="button"
      tabindex="0"
      on:click={() => {
        authHandlers.resetPassword(email);
      }}
      on:keydown={() => {}}
    >
      Forgot Password?
    </div>
  {/if}

  {#if displayError}
    <div>{errorMessage}</div>
  {/if}
</div>

<style>
  .container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    flex: 1;
  }

  .container form {
    display: flex;
    flex-direction: column;
  }

  /* make a div with red border */
  .container div {
    border: 1px solid red;
    padding: 10px;
    margin: 10px;
    cursor: pointer;
  }
</style>
