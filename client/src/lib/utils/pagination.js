/**
 * Get pagination information
 * @param {number} total - Total items
 * @param {number} limit - Items per page
 * @param {number} offset - Current offset
 */
export const getPagination = (total, limit, offset) => {
	const pages = Math.ceil(total / limit);
	const current = Math.ceil(offset / limit) + 1;
	const prev = current > 1 ? current - 1 : null;
	const next = current < pages ? current + 1 : null;
	const visiblePages = [];

	// Add visible page numbers
	for (let i = 1; i <= pages; i++) {
		if (
			i === 1 || // First page
			i === pages || // Last page
			(i >= current - 2 && i <= current + 2) // Pages around current
		) {
			visiblePages.push(i);
		} else if (visiblePages[visiblePages.length - 1] !== null) {
			visiblePages.push(null); // Add ellipsis
		}
	}

	return {
		pages,
		current,
		prev,
		next,
		offset,
		limit,
		visiblePages: visiblePages.filter(
			(page, index, array) =>
				page !== null || (array[index - 1] !== null && array[index + 1] !== null)
		)
	};
};
