<script>
  import { createLink } from '$lib/repositories/links'
  import { urlByShortId, urlWithoutProtocol } from '$lib/helpers/link'
  import { onMount } from 'svelte'
  import { fade } from 'svelte/transition';

  onMount(async () => {
    displayedOptions = false;
    inputUrl.focus();
  });

  async function shortenLink() {
    shortenLinkButtonCaption = '<span class="loading loading-dots loading-xs"></span>';
    let response = await createLink(url, password);
    let data = await response.json();

    if (response.status === 201) {
      shortenedUrl = urlByShortId(data.id);
    } else {
      shortenLinkButtonCaption = "Skrati!";
      if (data.description) {
        errorDescription = data.description;
      } else {
        errorDescription = "Došlo je do greške prilikom skraćivanja linka."
      }
    }
  }

  async function handleInput() {
    displayedOptions = true;
  }

  let displayedOptions;

  let url;
  let password;
  let inputUrl;
  let errorDescription;
  let shortenedUrl;
  let shortenLinkButtonCaption = "Skrati!";
</script>

<div class="text-center">
  {#if shortenedUrl === undefined}
    <div class="flex flex-col md:flex-row">
      <div class="flex-1 md:mr-4">
        <label class="input input-bordered flex items-center gap-2">
          🔗
          <input bind:this={inputUrl} bind:value={url} on:input={handleInput} id="url" type="text" class="w-full text-dark-blue" placeholder="Zalijepite link koji želite skratiti..." />
        </label>
      </div>
      <div class="mt-4 md:mt-0">
        <button on:click={shortenLink} class="btn btn-small btn-primary w-full md:w-auto">{@html shortenLinkButtonCaption}</button>
      </div>
    </div>
    {#if displayedOptions}
      <div transition:fade class="fade-in mt-4">
        <div class="flex flex-col md:flex-row">
          <div class="flex-1">
            <label class="input input-bordered flex items-center gap-2">
              🔒
              <input bind:value={password} type="password" maxlength="72" class="w-full text-dark-blue" placeholder="Ako želite, zaštitite skraćeni link lozinkom..." />
            </label>
          </div>
          <!-- div class="flex-1 mt-4 md:mt-0">
            <select class="select select-bordered w-full text-dark-blue">
              <option disabled selected class="text-yellow">🕙 Odaberite trajanje</option>
              <option>Trajanje1</option>
              <option>Trajanje2</option>
            </select>
          </div -->
        </div>
      </div>
      {#if errorDescription}
        <div class="mt-8 mb-4">
          <p class="text-sm">Došlo je do greške prilikom skraćivanja linka:</p>
          <p class="text-red-300">{errorDescription}</p>
        </div>
      {/if}
    {/if}
  {:else}
    <div class="mt-8 mb-8">
      <p class="text-2xl font-bold">
        <a href="{shortenedUrl}" target="_blank" class="text-blue">{urlWithoutProtocol(shortenedUrl)}</a>
        <span class="ml-2" style="font-size: .875em; margin-right: .125em; position: relative; top: -.25em; left: -.125em">
          📄
          <span style="position: absolute; top: .25em; left: .25em">📄</span>
        </span>
      </p>
    </div>
  {/if}
</div>

<style>
  .fade-in {
    transition: opacity 0.5s;
  }
</style>