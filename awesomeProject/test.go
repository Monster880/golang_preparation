package main

//We want to implement a meeting scheduler. It books time slots for incoming meetings.
//	A meeting occupies one or more consecutive time slots, and never overlaps any other meetings.
//	Initially there are 8 available time slots: 0, 1, 2, ..., 7.
//
//The scheduler supports two operations:
//
//(A) Scheduling a meeting - Given the meeting duration D (an integer within 1..8), the scheduler books D consecutive available time slots,
//or rejects the meeting if there are no such time slots.
//
//(B) Canceling a meeting - Given a scheduled meeting, the scheduler frees up all time slots which were booked for this meeting.

type meeting_scheduler struct {
	books_time     []int
	meetings_index []byte
}

var meetings_scheduler1 = &meeting_scheduler{books_time: []int{0, 1, 2, 3, 4, 5, 6, 7}}

func (*meeting_scheduler) booking_meeting(meeting_duration int) bool {
	max_start, max_end := slide_window(meetings_scheduler1.books_time)
	if max_end-max_start >= meeting_duration {
		for i := max_start; i < len(meetings_scheduler1.books_time); i++ {
			meetings_scheduler1.books_time[i] = -1
		}
		return true
	} else {
		return false
	}
}

func slide_window(books_time []int) (start, end int) {
	temp_end := 0
	max_start := -1
	max_end := -1
	res := 0
	for start := 0; start < len(books_time); start++ {
		if books_time[start] > 0 {
			max_start = start
		}
		for end := start; end < len(books_time); end++ {
			if books_time[temp_end] > 0 {
				temp_end++
			}
		}
		length := temp_end - start
		if length > res {
			res = length
			max_start = start
			max_end = temp_end
		}
	}
	return max_start, max_end
}

func (*meeting_scheduler) canceling_meeting() {

}
