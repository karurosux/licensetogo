export const getFileUrl = (
	collectionId: string | undefined,
	recordId: string | undefined,
	filename: string | undefined
) => {
	return `/api/files/${collectionId}/${recordId}/${filename}`;
};
