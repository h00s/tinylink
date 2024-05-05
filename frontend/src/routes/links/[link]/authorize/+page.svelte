<script>
  import HeaderContainer from '$lib/components/HeaderContainer.svelte'
  import Authorize from '$lib/components/Authorize.svelte'
  import Privacy from '$lib/components/Privacy.svelte'
  import Footer from '$lib/components/Footer.svelte'
  import { urlWithoutProtocol } from '$lib/helpers/link'

  export let data;
</script>

<HeaderContainer>
  <h2 class="pt-4 text-2xl font-extrabold text-center">
    {urlWithoutProtocol(data.url)}
  </h2>
  <p class="text-center pb-8 text-light-green">je zaštićen lozinkom</p>
  <Authorize link={data.link} />
</HeaderContainer>

<Privacy />
<Footer />