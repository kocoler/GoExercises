/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        queue<TreeNode*> q;
        q.push(root);
        vector<vector<int> > res;
        if (root == NULL) return res;
        bool flag = true; // true: left, false: right
        while (!q.empty()) {
            int size = q.size();
            vector<int> temp;
            if (flag) {
                for (int i = 0; i < size; i ++) {
                    TreeNode* node = q.front();
                    temp.push_back(node->val);         
                    q.pop();
                    if (node->left != NULL) q.push(node->left);
                    if (node->right != NULL) q.push(node->right);
                }
            } else {
                for (int i = 0; i < size; i ++) {
                    TreeNode* node = q.front();
                    temp.insert(temp.begin(), node->val);
                    q.pop();
                    if (node->left != NULL) q.push(node->left);
                    if (node->right != NULL) q.push(node->right);
                }
            }
            flag = !flag;
            res.push_back(temp);
        }

        return res;
    }
};
