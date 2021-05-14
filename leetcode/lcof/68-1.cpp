/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if (q == p) {
            return q;
        }
        TreeNode* temp = root;
        TreeNode* prep = temp;
        while (temp != p) {
            if (temp == q) {
                return temp;
            }
            prep = temp;
            if (temp->val > p->val) {
                temp = temp->left;
            } else if (temp->val < p->val) {
                temp = temp->right;
            }
        }

        temp = root;
        TreeNode* preq = temp;
        while (temp != q) {
            if (temp == p) {
                return temp;
            }
            preq = temp;
            if (temp->val > q->val) {
                temp = temp->left;
            } else if (temp->val < q->val) {
                temp = temp->right;
            }
        }

        return lowestCommonAncestor(root, prep, preq);
    }
};
