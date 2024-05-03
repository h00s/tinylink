<script>
  import FlexContainer from '$lib/components/FlexContainer.svelte'
  import Authorize from '$lib/components/Authorize.svelte'
  import Privacy from '$lib/components/Privacy.svelte'
  import Footer from '$lib/components/Footer.svelte'

  export let data;
</script>

<FlexContainer>
  <h1 class="title pt-16 pb-2 font-extrabold text-center">skrati.link</h1>
  <h2 class="pt-4 text-2xl font-extrabold text-center">
    {data.url}
  </h2>
  <p class="text-center pb-8 text-light-green">je zaštićen lozinkom.</p>
  <Authorize link={data.link} />
</FlexContainer>

<Privacy />
<Footer />