<script>
  import { createLink } from '$lib/repositories/links'
  import { urlByShortId, linkModel } from '$lib/helpers/link'
  import { onMount } from 'svelte'
  import { fade } from 'svelte/transition';

  onMount(async () => {
    displayedOptions = false;
    inputUrl.focus();
  });

  async function shortenLink() {
    let response = await createLink(url);
    let data = await response.json();

    if (response.status === 201) {
      shortenedUrl = urlByShortId(data.id);
    } else {
      if (data.description) {
        alert(data.description);
      } else {
        alert('An error occurred while shortening the link');
      }
    }
  }

  async function handleInput() {
    displayedOptions = true;
  }

  let displayedOptions;

  let url;
  let inputUrl;
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
        <button on:click={shortenLink} class="btn btn-small btn-primary w-full md:w-auto">{shortenLinkButtonCaption}</button>
      </div>
    </div>
    {#if displayedOptions}
      <div transition:fade class="fade-in mt-4">
        <div class="flex flex-col md:flex-row">
          <div class="flex-1">
            <label class="input input-bordered flex items-center gap-2">
              🔒
              <input id="password" type="password" class="w-full text-dark-blue" placeholder="Ako želite, zaštitite skraćeni link lozinkom..." />
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
    {/if}
  {:else}
    <div class="mb-4">
      <p>
        Your shortened link is:
        {shortenedUrl}
      </p>
    </div>
  {/if}
</div>

<style>
  .fade-in {
    transition: opacity 0.5s;
  }
</style>