package queue

type Queue struct {
	Length int
	Front  *ListNode
	Rear   *ListNode
}

type ListNode struct {
	Data string
	Next *ListNode
}

func (q *Queue) length() int {
	return q.Length
}

func (q *Queue) isEmpty() bool {
	return q.Length == 0
}

func (q *Queue) offer(data string) {
	temp := &ListNode{
		Data: data,
	}

	if q.isEmpty() {
		q.Front = temp
	} else {
		q.Rear.Next = temp
	}
	q.Rear = temp
	q.Length++
}

func (q *Queue) poll() string {
	if q.isEmpty() {
		return ""
	}

	temp := q.Front
	q.Front = q.Front.Next
	q.Length--

	return temp.Data
}

func (q *Queue) generateBinaryNumbers(n int) []string {
	result := make([]string, n)
	q.offer("1")

	for i := 0; i < n; i++ {
		result[i] = q.poll()
		n1 := result[i] + "0"
		n2 := result[i] + "1"
		q.offer(n1)
		q.offer(n2)
	}

	return result
}
