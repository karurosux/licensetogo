export const GET = async ({ params, locals }) => {
	const { collectionId, filename, recordId } = params;
	const url = locals.pb.files.getURL({ id: recordId, collectionId }, filename);

	try {
		// Fetch the image from the URL
		const response = await fetch(url);

		if (!response.ok) {
			return new Response('Failed to fetch the image', { status: response.status });
		}

		// Get the content type from the response
		const contentType = response.headers.get('content-type') || 'application/octet-stream';

		// Get the image data as an ArrayBuffer
		const imageData = await response.arrayBuffer();

		// Return the image with the appropriate content type
		return new Response(imageData, {
			headers: {
				'Content-Type': contentType,
				'Cache-Control': 'public, max-age=3600'
			}
		});
	} catch (error) {
		console.error('Error fetching image:', error);
		return new Response('Error fetching the image', { status: 500 });
	}
};
