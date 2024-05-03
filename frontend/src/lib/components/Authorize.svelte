<script>
  import { authLink } from '$lib/repositories/links'
  import { onMount } from 'svelte'

  onMount(async () => {
    inputPassword.focus();
  });

  async function authorize() {
    let response = await authLink(link, password);
    let data = await response.json();

    if (response.status === 200) {
      window.location = data.url;
    } else {
      password = '';
      inputPassword.focus();
      error = data.description;
    }
  }

  let error;
  let inputPassword;
  let password;

  export let link;
</script>

<div class="text-center">
  <div class="flex flex-col md:flex-row">
    <div class="flex-1 md:mr-4">
      <label class="input input-bordered flex items-center gap-2">
        🔒
        <input bind:this={inputPassword} bind:value={password} type="password" class="w-full text-dark-blue" />
      </label>
    </div>
    <div class="mt-4 md:mt-0">
      <button on:click={authorize} class="btn btn-small btn-primary w-full md:w-auto">Autoriziraj</button>
    </div>
  </div>
  {#if error}
    <div class="mb-4">
      <p>
        {error}
      </p>
    </div>
  {/if}
</div>