import { fetchLink } from '$lib/repositories/links';
import { redirect } from '@sveltejs/kit';

export async function load({ fetch, params }) {
  const link = await fetchLink(params.link, fetch);
  return redirect(301, link.url);

  /*return {
    link: params.link,
  };*/
}