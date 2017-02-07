function accumulate(input, fn)
	local output = {}
	for _, v in ipairs(input) do
		table.insert(output, fn(v))
	end
	return output
end

return accumulate
