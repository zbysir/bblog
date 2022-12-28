package hollow

import (
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/stretchr/testify/assert"
	jsx "github.com/zbysir/gojsx"
	"github.com/zbysir/hollow/internal/pkg/gobilly"
	"testing"
)

func TestJsx(t *testing.T) {
	jx, err := jsx.NewJsx(jsx.Option{
		SourceCache: nil,
		Debug:       false,
	})
	if err != nil {
		t.Fatal(err)
		return
	}
	f := memfs.New()
	file, _ := f.Create("index.tsx")
	file.Write([]byte(`export default function HelloJSX({name}){ return <>{name}</> }`))

	m := NewGoldMdRender(GoldMdRenderOptions{
		jsx:              jx,
		fs:               gobilly.NewStdFs(f),
		assetsUrlProcess: nil,
	})

	cases := []struct {
		Name string
		In   []byte
		Get  []byte
	}{
		{
			Name: "Jsx",
			In: []byte(`
import Name from './index.tsx'

Hello

<Name name={'bysir'}>
</Name>
`),
			Get: []byte(`<p>Hello</p>
bysir`),
		},
		{
			Name: "Empty(Currently not supported)",
			In: []byte(`
import Name from './index.tsx'

Hello

<>
</>

<Name name={'bysir'}>
</Name>
`),
			Get: []byte("<p>Hello</p>\n<p>&lt;&gt;\n&lt;/&gt;</p>\nbysir"),
		},
		{
			Name: "Jsx Inline",
			In: []byte(`
import Name from './index.tsx'

Hello

<Name name={'bysir'}></Name>

!`),
			Get: []byte(`<p>Hello</p>
bysir<p>!</p>
`),
		},
		{
			Name: "Jsx SelfClose",
			In: []byte(`
import Name from './index.tsx'

Hello

<Name name={'bysir'} /> 333

!`),
			Get: []byte(`<p>Hello</p>
bysir<p>!</p>
`),
		},
		{
			Name: "Jsx with md",
			In: []byte(`
import Name from './index.tsx'

Hello

<Name name={'bysir'}>

**strong**

</Name>
`),
			Get: []byte(`<p>Hello</p>
bysir`),
		},
		{
			Name: "Base",
			In: []byte(`
Hello

<Warning>

**strong**

</Warning>
`),
			Get: []byte(`<p>Hello</p>
<Warning>
<p><strong>strong</strong></p>
</Warning>
`),
		},
		{
			Name: "WithMeta",

			In: []byte(`---
a: 1
---

import Name from './index.tsx'

Hello

<Name name={'bysir'}>
</Name>

!
`),
			Get: []byte(`<p>Hello</p>
bysir<p>!</p>
`),
		}, {
			Name: "A",

			In: []byte(`---
a: 1
---

import A from './index.tsx'

Hello

<A name={'bysir'}>
</A>

!
`),
			Get: []byte(`<p>Hello</p>
bysir<p>!</p>
`),
		},
		{
			Name: "WithOutImport",

			In: []byte(`---
a: 1
---
Hello

<Name name={'bysir'}>
</Name>
`),
			Get: []byte(`<p>Hello</p>
<p>&lt;Name name={'bysir'}&gt;
</Name></p>
`),
		},
		{
			Name: "Empty",

			In: []byte(`
<>
	{[1,2].map((i)=><p> {i} </p>)}
</>
`),
			Get: []byte(`<p> 1 </p><p> 2 </p>`),
		},
		{
			Name: "Inline",

			In: []byte(`
const A = ({name}) => <>Hello {name}</>
const name = 1

<A name={name}></A>
`),
			Get: []byte(`Hello 1`),
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			s, err := m.Render(c.In)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, string(c.Get), string(s.Body))
		})
	}

}

func TestAssert(t *testing.T) {
	jx, err := jsx.NewJsx(jsx.Option{
		SourceCache: nil,
		Debug:       false,
	})
	if err != nil {
		t.Fatal(err)
		return
	}
	f := memfs.New()
	file, _ := f.Create("index.tsx")
	file.Write([]byte(`export default function HelloJSX({name}){ return <>{name}</> }`))

	m := NewGoldMdRender(GoldMdRenderOptions{
		jsx: jx,
		fs:  gobilly.NewStdFs(f),
		assetsUrlProcess: func(s string) string {
			return "1"
		},
	})

	a, _ := m.Render([]byte(`## h2
![在 Golang 中尝试“干净架构”](../../static/img/在%20Golang%20中尝试干净架构_1.png)
![在 Golang 中尝试“干净架构”](/static/img/在%20Golang%20中尝试干净架构_1.png)

在[文中](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)，他提出的干净架构是这样的：

<HelloJSX></HelloJSX>
`))
	t.Logf("%s", a.Body)
}
