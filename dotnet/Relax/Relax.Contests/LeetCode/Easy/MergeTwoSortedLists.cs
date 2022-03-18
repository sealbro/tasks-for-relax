using System;

namespace Relax.Contests.LeetCode.Easy
{
    /// <summary>
    /// https://leetcode.com/problems/merge-two-sorted-lists/
    /// </summary>
    public static class MergeTwoSortedLists
    {
        public static void MainX(string[] args)
        {
            var list1 = new[] { 1, 2, 4 };
            var list2 = new[] { 1, 3, 4 };
            // var list1 = new[] { -9, 3 };
            // var list2 = new[] { 5, 7 };

            var solution = new Solution();
            var mergeTwoLists = solution.MergeTwoLists(FillNodes(list1), FillNodes(list2));

            ListNode node = mergeTwoLists;
            while (node != null)
            {
                Console.WriteLine(node.val);
                node = node.next;
            }
        }

        private static ListNode FillNodes(int[] ints)
        {
            ListNode first = null;
            ListNode current = null;

            foreach (var val in ints)
            {
                if (first == null)
                {
                    first = new ListNode(val);
                    current = first;
                }
                else
                {
                    current.next = new ListNode(val);
                    current = current.next;
                }
            }

            return first;
        }
    }

    public class ListNode
    {
        public int val;
        public ListNode next;

        public ListNode(int val = 0, ListNode next = null)
        {
            this.val = val;
            this.next = next;
        }
    }

    public class Solution
    {
        public ListNode MergeTwoLists(ListNode list1, ListNode list2)
        {
            ListNode result = null;
            ListNode currentNode = null;
            while (list1 != null || list2 != null)
            {
                ListNode current;
                if (list1 == null)
                {
                    if (result == null)
                    {
                        result = list2;
                    }
                    else
                    {
                        currentNode.next = list2;
                    }
                    break;
                }

                if (list2 == null)
                {
                    if (result == null)
                    {
                        result = list1;
                    }
                    else
                    {
                        currentNode.next = list1;
                    }
                    break;
                }

                if (list1.val > list2.val)
                {
                    current = list2;
                    list2 = list2.next;
                }
                else
                {
                    current = list1;
                    list1 = list1.next;
                }

                if (result == null)
                {
                    result = current;
                    currentNode = result;
                }
                else
                {
                    currentNode.next = current;
                    currentNode = currentNode.next;
                }
            }

            return result;
        }
    }
}