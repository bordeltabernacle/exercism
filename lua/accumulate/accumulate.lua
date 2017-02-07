return function(input, fn)
	local output = {}
	for index, value in pairs(input) do
		output[index] = fn(value)
	end
	return output
end
