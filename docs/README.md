<p>Packages:</p>
<ul>
<li>
<a href="#fpetkovski.io%2fv1alpha1">fpetkovski.io/v1alpha1</a>
</li>
</ul>
<h2 id="fpetkovski.io/v1alpha1">fpetkovski.io/v1alpha1</h2>
Resource Types:
<ul></ul>
<h3 id="fpetkovski.io/v1alpha1.ResourceRule">ResourceRule
</h3>
<p>
(<em>Appears on:</em><a href="#fpetkovski.io/v1alpha1.TTLPolicySpec">TTLPolicySpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="fpetkovski.io/v1alpha1.TTLPolicy">TTLPolicy
</h3>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#fpetkovski.io/v1alpha1.TTLPolicySpec">
TTLPolicySpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#fpetkovski.io/v1alpha1.ResourceRule">
ResourceRule
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ttlFrom</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>expirationFrom</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#fpetkovski.io/v1alpha1.TTLPolicyStatus">
TTLPolicyStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="fpetkovski.io/v1alpha1.TTLPolicySpec">TTLPolicySpec
</h3>
<p>
(<em>Appears on:</em><a href="#fpetkovski.io/v1alpha1.TTLPolicy">TTLPolicy</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#fpetkovski.io/v1alpha1.ResourceRule">
ResourceRule
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ttlFrom</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>expirationFrom</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="fpetkovski.io/v1alpha1.TTLPolicyStatus">TTLPolicyStatus
</h3>
<p>
(<em>Appears on:</em><a href="#fpetkovski.io/v1alpha1.TTLPolicy">TTLPolicy</a>)
</p>
<p>
</p>
<hr/>
<p><em>
Generated with <code>gen-crd-api-reference-docs</code>
on git commit <code>f52dd5a</code>.
</em></p>
