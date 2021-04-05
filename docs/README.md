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
<p>ResourceRule defines the resources to which the TTLPolicy should be applied</p>
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
<p>APIVersion is the full API version of the kubernetes resources. <br />
Examples: <br />
- v1 <br />
- apps/v1 <br /></p>
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
<p>Kind is the resources&rsquo; Kind.
Examples: <br />
- Deployment <br />
- Ingress <br /></p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Namespace is the namespace in which the resources are created</p>
</td>
</tr>
<tr>
<td>
<code>matchLabels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>MatchLabels is the label set which the resources should match</p>
</td>
</tr>
</tbody>
</table>
<h3 id="fpetkovski.io/v1alpha1.TTLPolicy">TTLPolicy
</h3>
<p>
<p>TTLPolicy is the object through which time to live behavior is configured for a Kubernetes resource.</p>
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
<p>TTLPolicySpec is the spec of the TTLPolicy</p>
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
<p>ResourceRule defines the resources to which the TTLPolicy should be applied</p>
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
<p>TTLFrom is the resources&rsquo; property which contains the TTL value for the specific resource. <br/>
Examples: <br />
- 15s <br />
- 1m <br />
- 1h30m <br /></p>
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
<em>(Optional)</em>
<p>ExpirationFrom is the resources&rsquo; property which contains the time from which TTL is calculated.
Examples include <code>.metadata.creationTimestamp</code> or <code>.status.startTime</code>.
The time should be specified in in <code>RFC3339</code> format.</p>
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
<p>TTLPolicySpec is the status of the TTLPolicy</p>
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
<p>ResourceRule defines the resources to which the TTLPolicy should be applied</p>
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
<p>TTLFrom is the resources&rsquo; property which contains the TTL value for the specific resource. <br/>
Examples: <br />
- 15s <br />
- 1m <br />
- 1h30m <br /></p>
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
<em>(Optional)</em>
<p>ExpirationFrom is the resources&rsquo; property which contains the time from which TTL is calculated.
Examples include <code>.metadata.creationTimestamp</code> or <code>.status.startTime</code>.
The time should be specified in in <code>RFC3339</code> format.</p>
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
on git commit <code>41406e2</code>.
</em></p>
