<script>
  import { authLink } from '$lib/repositories/links'
  import { tick } from 'svelte'
  import { afterNavigate } from '$app/navigation'

  afterNavigate(async () => {
    await tick();
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
      errorDescription = data.description;
    }
  }

  let errorDescription;
  let password, inputPassword;

  export let link;
</script>

<div class="text-center">
  <div class="flex flex-col md:flex-row">
    <div class="flex-1 md:mr-4">
      <label class="input input-bordered flex items-center gap-2">
        🔒
        <input bind:this={inputPassword} bind:value={password} type="password" class="w-full text-dark-blue" placeholder="Upišite lozinku da biste otključali skraćeni link..." />
      </label>
    </div>
    <div class="mt-4 md:mt-0">
      <button on:click={authorize} class="btn btn-small btn-primary w-full md:w-auto">Otključaj</button>
    </div>
  </div>
  {#if errorDescription}
    <div class="mt-8 mb-4">
      <p class="text-red-300">{errorDescription}</p>
    </div>
  {/if}
</div>