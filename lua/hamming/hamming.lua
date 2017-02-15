local Hamming = {}

function Hamming.compute(a, b)
	local d = 0
	if #a ~= #b then return -1 end
	for i = 1, #a do
		if a:byte(i) ~= b:byte(i) then
			d = d + 1
		end
	end
	return d
end

return Hamming
