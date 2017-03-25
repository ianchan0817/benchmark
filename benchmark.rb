current_time = Time.now
puts "********* Bechmark start *********"
largest_number = 0

for x in 1..10000000
  largest_number = x if largest_number < x
end
duration = Time.now - current_time
puts "It takes #{duration*100} ms\n**********************************\n"


# ********* Bechmark start *********
# It takes 56.7782 ms
# **********************************


